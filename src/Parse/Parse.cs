using System.Collections.Generic;
using System.Text;



namespace TypeGo
{

    public class ParseData
    {
        public List<Token> TokenList = new();
        public Token? LastToken = null;
        public bool WasComment = false;
        public int CharacterIndex = 0;
        public string Code = "";
        public int LineCount = 0; //for token position
        public int CharCount = 0;
        public ParseResult ParseResult = ParseResult.Ok;
    }

    public static class Parse
    {


        static bool ShouldSkip(ref ParseData parseData)
        {
            parseData.CharCount += 1;
            if (parseData.LastToken.HasValue) {

                if (parseData.LastToken.Value.Type == TokenType.Comment) {
                    parseData.WasComment = true;
                }
            }

            char currentChar = parseData.Code[parseData.CharacterIndex];

            if (currentChar == '\n') {

                if (parseData.WasComment == true) {
                    parseData.TokenList.Add(new Token("\\n", TokenType.EndComment, parseData.LineCount, parseData.CharCount));
                    parseData.WasComment = false;
                } else {
                    parseData.TokenList.Add(new Token("\\n", TokenType.NewLine, parseData.LineCount, parseData.CharCount));
                }
                parseData.LineCount += 1;
                parseData.CharCount = 0;
                parseData.CharacterIndex += 1;
                return true;
            }
            bool isSpecialChar =
                currentChar == '\r' ||
                currentChar == '\t' ||
                currentChar == ' ' ||
                currentChar == '\\';

            if (isSpecialChar == true) {
                if (currentChar == '\t') {
                    Token tabToken = new Token("\\t", TokenType.Tab, parseData.LineCount, parseData.CharCount);
                }
                parseData.CharacterIndex += 1;
                return true;
            }
            return false;
        }

        static bool IsInvalid(Token? token, ref ParseData parseData)
        {
            if (token.HasValue == false)
            {
                return true;
            }
            if (parseData.ParseResult != ParseResult.Ok)
            {
                return true;
            }
            if (token.Value.Text == null)
            {
                return true;
            }
            return false;
        }

        static void ProcessCharacter(ref ParseData parseData)
        {
            if (ShouldSkip(ref parseData) == true) return;

            int previousCharacterIndex = parseData.CharacterIndex;

            Token? token = GetToken(ref parseData);

            if (previousCharacterIndex == parseData.CharacterIndex) {
                parseData.CharacterIndex += 1;
            }
            if (IsInvalid(token, ref parseData) == true) {
                return;
            }

            parseData.TokenList.Add(token.Value);
            parseData.LastToken = token;
        }

        public static List<Token> ParseToTokens(ref ParseResult error, string code)
        {
            //Fmt.PrintColor("\tParsing\t\t\t\t", ConsoleColor.DarkGray);

            ParseData parseData = new();
            parseData.Code = code;

            if (parseData.Code.Length == 0) {
                error = ParseResult.String_Length_Zero;
                return parseData.TokenList;
            }

            int STRING_LENGTH = code.Length;
            for (parseData.CharacterIndex = 0; parseData.CharacterIndex < STRING_LENGTH; ) {

                ProcessCharacter(ref parseData);
                if (parseData.ParseResult != ParseResult.Ok) {
                    error = parseData.ParseResult;
                    return parseData.TokenList;
                }

            }

            return parseData.TokenList;
        }
        static Token? GetToken(ref ParseData parseData)
        {

            char currentChar = parseData.Code[parseData.CharacterIndex];

            if (currentChar == '"')
            {
                return ReadString(ref parseData);
            }
            if (currentChar == '\'')
            {
                return ReadChar(ref parseData);
            }
            if (ParseUtils.IsOperator(currentChar))
            {
                return ReadOperator(ref parseData);
            }
            if (ParseUtils.IsSeparator(currentChar))
            {
                if (currentChar == '\n')
                {
                    parseData.WasComment = false;
                }

                return ReadSeparator(ref parseData);
            }

            return ReadWord(ref parseData);
        }

        static Token? ReadString(ref ParseData parseData)
        {
            StringBuilder sb = new();
            sb.Append(parseData.Code[parseData.CharacterIndex]);
            parseData.CharacterIndex += 1;

            while (parseData.CharacterIndex < parseData.Code.Length)
            {
                char c = parseData.Code[parseData.CharacterIndex];
                if (c == '"')
                {
                    sb.Append(c);
                    parseData.CharacterIndex++;
                    return new Token(sb.ToString(), TokenType.StringValue, parseData.LineCount, parseData.CharCount);
                }
                sb.Append(c);
                parseData.CharacterIndex++;
            }

            parseData.ParseResult = ParseResult.Unterminated_String;
            return null;
        }

        static Token? ReadSeparator(ref ParseData parseData)
        {
            char c = parseData.Code[parseData.CharacterIndex];
            parseData.CharacterIndex += 1;
            string tokenText = c.ToString();
            return new Token(tokenText, ParseUtils.GetTokenType(parseData, tokenText), parseData.LineCount, parseData.CharCount);
        }

        static Token? ReadChar(ref ParseData parseData)
        {
            if (parseData.WasComment) {
                return new Token("'", TokenType.Comma, parseData.LineCount, parseData.CharCount);
            }

            int start = parseData.CharacterIndex;
            parseData.CharacterIndex += 1;
            if (parseData.CharacterIndex >= parseData.Code.Length)
            {
                parseData.ParseResult = ParseResult.Unexpected_Value;
                return null;
            }

            char charValue = parseData.Code[parseData.CharacterIndex];
            char? charValueSecondPart = null;
            parseData.CharacterIndex += 1;

            if (charValue == '\\') {
                charValueSecondPart = parseData.Code[parseData.CharacterIndex];
                parseData.CharacterIndex += 1;
            }

            if (parseData.CharacterIndex >= parseData.Code.Length) {
                parseData.ParseResult = ParseResult.Unexpected_Value;
                return null;
            }
            if (parseData.Code[parseData.CharacterIndex] != '\'') {
                parseData.ParseResult = ParseResult.Unterminated_Char;
                return null;
            }
            parseData.CharacterIndex += 1;

            if (charValueSecondPart != null) {
                return new Token($"'{charValue}{charValueSecondPart}'", TokenType.CharValue, parseData.LineCount, parseData.CharCount);
            }
            return new Token($"'{charValue}'", TokenType.CharValue, parseData.LineCount, parseData.CharCount);
        }

        static Token? ReadOperator(ref ParseData parseData)
        {
            StringBuilder sb = new();
            char c = parseData.Code[parseData.CharacterIndex];
            sb.Append(c);

            parseData.CharacterIndex += 1;

            // Lookahead for compound operators like "==", "!="
            if (parseData.CharacterIndex < parseData.Code.Length) {

                char next = parseData.Code[parseData.CharacterIndex]; 
                if (ParseUtils.IsOperator(next)) {
                    sb.Append(next);
                    parseData.CharacterIndex += 1;
                }
            }

            string tokenText = sb.ToString();
            return new Token(tokenText, ParseUtils.GetTokenType(parseData, tokenText), parseData.LineCount, parseData.CharCount);
        }

        static Token? ReadWord(ref ParseData parseData)
        {
            StringBuilder sb = new();

            while (parseData.CharacterIndex < parseData.Code.Length)
            {
                char c = parseData.Code[parseData.CharacterIndex];

                if (char.IsLetterOrDigit(c) || c == '_') {
                    sb.Append(c);
                    parseData.CharacterIndex++;
                } else {
                    break;
                }
            }

            string word = sb.ToString();
            return new Token(word, ParseUtils.GetTokenType(parseData, word), parseData.LineCount, parseData.CharCount);
        }
    }

    public static class ParseUtils
    {
        public static bool IsOperator(char c)
        {
            for (int i = 0; i < SYNTAX.OPERATORS.Length; i++) {

                if (c == SYNTAX.OPERATORS[i]) {
                    return true;
                }
            }
            return false;
        }

        public static bool IsSeparator(char c)
        {
            for (int i = 0; i < SYNTAX.SEPERATORS.Length; i++)
            {

                if (c == SYNTAX.SEPERATORS[i])
                {
                    return true;
                }
            }
            return false;
        }

        public static bool IsInteger(string input)
        {
            return int.TryParse(input, out _);
        }
        public static bool IsDecimal(string input)
        {
            return double.TryParse(input, out _);
        }

        public static TokenType GetTokenType(ParseData parseData, string input)
        {
            // Keywords
            if (input == Constants.FN) return TokenType.Fn;
            if (input == Constants.FUNC) return TokenType.Fn;
            if (input == Constants.IF) return TokenType.If;
            if (input == Constants.ELSE) return TokenType.Else;
            if (input == Constants.FOR) return TokenType.For;
            if (input == Constants.WHILE) return TokenType.While;
            if (input == Constants.RETURN) return TokenType.Return;
            if (input == Constants.BREAK) return TokenType.Break;
            if (input == Constants.CONTINUE) return TokenType.Continue;
            if (input == Constants.PRINT) return TokenType.Print;
            if (input == Constants.PRINTLN) return TokenType.Println;
            if (input == Constants.TRUE) return TokenType.True;
            if (input == Constants.FALSE) return TokenType.False;

            // Types
            if (input == Constants.INT) return TokenType.Int;
            if (input == Constants.INT8) return TokenType.Int8;
            if (input == Constants.INT16) return TokenType.Int16;
            if (input == Constants.INT32) return TokenType.Int32;
            if (input == Constants.INT64) return TokenType.Int64;
            if (input == Constants.UINT16) return TokenType.Uint16;
            if (input == Constants.UINT32) return TokenType.Uint32;
            if (input == Constants.UINT64) return TokenType.Uint64;
            if (input == Constants.F32) return TokenType.Float32;
            if (input == Constants.F64) return TokenType.Float64;
            if (input == Constants.STRING) return TokenType.String;
            if (input == Constants.BOOL) return TokenType.Bool;
            if (input == Constants.BYTE) return TokenType.Byte;
            if (input == Constants.RUNE) return TokenType.Rune;
            if (input == Constants.CONST) return TokenType.Const;
            if (input == Constants.STRUCT) return TokenType.Struct;
            if (input == Constants.ENUM) return TokenType.Enum;
            if (input == Constants.ENUMSTRUCT) return TokenType.Enumstruct;
            if (input == Constants.MAP) return TokenType.Map;
            if (input == Constants.ERROR) return TokenType.Error;
            if (input == Constants.DEFER) return TokenType.Defer;
            if (input == Constants.CHAN) return TokenType.Chan;
            if (input == Constants.CHANNEL_SETTER) return TokenType.Channel_Setter;
            if (input == Constants.VAR) return TokenType.Var;
            if (input == Constants.GO) return TokenType.Go;
            if (input == Constants.TYPE) return TokenType.Type;
            if (input == Constants.INTERFACE) return TokenType.Interface;
            if (input == Constants.GOTO) return TokenType.Goto;
            if (input == Constants.MINUS_MINUS) return TokenType.MinusMinus;
            if (input == Constants.NOT) return TokenType.Not;

            // Operators
            if (input == Constants.PLUS_PLUS) return TokenType.PlusPlus;
            if (input == Constants.PLUS) return TokenType.Plus;
            if (input == Constants.MINUS) return TokenType.Minus;
            if (input == Constants.MULTIPLY) return TokenType.Multiply;
            if (input == Constants.DIVIDE) return TokenType.Divide;
            if (input == Constants.EQUALS) return TokenType.Equals;
            if (input == Constants.PLUS_EQUALS) return TokenType.PlusEquals;
            if (input == Constants.MINUS_EQUALS) return TokenType.MinusEquals;
            if (input == Constants.MULTIPLY_EQUALS) return TokenType.MultiplyEquals;
            if (input == Constants.DIVIDE_EQUALS) return TokenType.DivideEquals;
            if (input == Constants.GREATER_THAN) return TokenType.GreaterThan;
            if (input == Constants.LESS_THAN) return TokenType.LessThan;
            if (input == Constants.EQUALS_EQUALS) return TokenType.EqualsEquals;
            if (input == Constants.GREATER_THAN_EQUALS) return TokenType.GreaterThanEquals;
            if (input == Constants.LESS_THAN_EQUALS) return TokenType.LessThanEquals;
            if (input == Constants.MODULUS) return TokenType.Modulus;
            if (input == Constants.NOT_EQUALS) return TokenType.NotEquals;
            if (input == Constants.AND) return TokenType.And;
            if (input == Constants.AND_AND) return TokenType.AndAnd;
            if (input == Constants.OR) return TokenType.Or;
            if (input == Constants.OR_OR) return TokenType.OrOr;
            if (input == Constants.MODULUS_EQUALS) return TokenType.ModulusEquals;
            if (input == Constants.MULTILINE_COMMENT_START) return TokenType.MultiLineStart;
            if (input == Constants.MULTILINE_COMMENT_END) return TokenType.MultiLineEnd;

            if (input == Constants.COMMENT) {
                parseData.WasComment = true;
                return TokenType.Comment;
            }

            // Parentheses and Brackets
            if (input == Constants.LEFT_PARENTHESIS) return TokenType.LeftParenthesis;
            if (input == Constants.RIGHT_PARENTHESIS) return TokenType.RightParenthesis;
            if (input == Constants.LEFT_BRACE) return TokenType.LeftBrace;
            if (input == Constants.RIGHT_BRACE) return TokenType.RightBrace;
            if (input == Constants.LEFT_SQUARE_BRACKET) return TokenType.LeftSquareBracket;
            if (input == Constants.RIGHT_SQUARE_BRACKET) return TokenType.RightSquareBracket;

            if (input == Constants.COLON) return TokenType.Colon;
            if (input == Constants.COLON_EQUALS) return TokenType.ColonEquals;
            if (input == Constants.SEMICOLON) return TokenType.Semicolon;
            if (input == Constants.COMMA) return TokenType.Comma;
            if (input == Constants.FULL_STOP) return TokenType.FullStop;

            if (input == Constants.MAKE) return TokenType.Make;
            if (input == Constants.APPEND) return TokenType.Append;
            if (input == Constants.PACKAGE) return TokenType.Package;
            if (input == Constants.IMPORT) return TokenType.Import;
            if (input == Constants.SWITCH) return TokenType.Switch;

            // Number literals
            if (IsInteger(input)) return TokenType.IntegerValue;
            if (IsDecimal(input)) return TokenType.DecimalValue;

            // String or Char value
            if (input.Contains('"')) return TokenType.StringValue;
            if (input.Contains('\'')) return TokenType.CharValue;

            return TokenType.Identifier;
        }


        static PotentialType GetPotentialType(char c)
        {
            if (c == ' ')
            {
                return PotentialType.Seperator;
            }

            for (int i = 0; i < SYNTAX.SEPERATORS.Length; i++)
            {

                if (c == SYNTAX.SEPERATORS[i])
                {
                    return PotentialType.SeperatorToken;
                }
            }

            for (int i = 0; i < SYNTAX.OPERATORS.Length; i++)
            {

                if (c == SYNTAX.OPERATORS[i])
                {
                    return PotentialType.Operator;
                }
            }

            if (c == '\'')
            {
                return PotentialType.CharValue;
            }
            if (c == '"')
            {
                return PotentialType.String;
            }
            if (c == '*')
            {
                return PotentialType.Multiply;
            }
            return PotentialType.NA;
        }

        static CharacterAction GetAction(PotentialType potentialType, bool isString, bool isChar)
        {
            if (isString == true) {
                return CharacterAction.AddChar;
            }
            if (isChar == true) {
                return CharacterAction.AddChar;
            }

            if (potentialType == PotentialType.NA) {
                return CharacterAction.AddChar;
            }

            if (potentialType == PotentialType.Operator) {

                return CharacterAction.Operator;
            }

            if (potentialType == PotentialType.SeperatorToken) {
                return CharacterAction.AddCharAndBreak;
            }

            if (potentialType == PotentialType.Seperator) {

                return CharacterAction.Break;
            }

            return CharacterAction.Break;
        }
    }
}
