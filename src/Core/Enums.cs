

namespace TypeGo
{

    public enum SpecialTokenCase
    {
        None,
        Comment,
        MultilineComment,
    }
    public enum WhileLoopAction
    {
        None,
        Continue,
        Break,
        Return,
        Error,
    }
    public enum NodeType
    {
        Invalid,
        Channel_Declaration,
        Channel_Declaration_With_Value,
        Interface_Declaration,
        Single_Declaration_With_Value,
        Single_Declaration_No_Value,
        Multiple_Declarations_No_Value,
        Multiple_Declarations_With_Value,
        Multiple_Declarations_Same_Type_No_Value,
        Multiple_Declarations_Same_Type_With_Value,
        Multiple_Declarations_One_Type_One_Set_Value,
        Constant_Global_Variable,
        Constant_Global_Variable_With_Type,
        Struct_Variable_Declaration,
        If_Statement_With_Declaration,
        If_Statement,
        Else_Statement,
        For_Loop,
        For_Loop_With_Declaration,
        Err_Return,
        Multi_Line_Import,
        Single_Import,
        Single_Import_With_Alias,
        NestedStruct,
        Struct_Declaration,
        Enum_Declaration,
        Enum_Variable,
        Enum_Variable_With_Value,
        Enum_Struct_Declaration,
        NewLine,
        Comment,
        Append,
        Package,
        Other,
    }

    public enum TokenType
    {
        Fn,
        //Literal
        IntegerValue,
        DecimalValue,
        CharValue,
        StringValue,
        True,
        False,
        Switch,
        Channel_Setter,
        ErrReturn,
        Case,

        Int,
        Int8,
        Int16,
        Int32,
        Int64,

        Uint,
        Uint8,
        Uint16,
        Uint32,
        Uint64,

        Float32,
        Float64,

        String,
        Byte,
        Rune,
        Bool,

        Struct,
        Enum,
        Enumstruct,
        Error,

        Map,
        Chan,
        Var,
        Type,
        Interface,
        Goto,

        //Keywords
        If,
        Else,
        For,
        While,
        Return,
        Break,
        Continue,
        Defer,
        Go,
        Not,

        //Operator
        Minus,
        Plus,
        PlusPlus,
        Divide,
        Multiply,
        Equals,
        And,
        AndAnd,
        Or,
        OrOr,
        PlusEquals,
        MinusEquals,
        MultiplyEquals,
        DivideEquals,
        GreaterThan,
        LessThan,
        EqualsEquals,
        GreaterThanEquals,
        LessThanEquals,
        Modulus,
        ModulusEquals,
        NotEquals,
        Colon,
        ColonEquals,
        MinusMinus,

        Identifier,

        Print,
        Println,

        //brackets
        LeftParenthesis,
        RightParenthesis,
        LeftBrace,
        RightBrace,
        LeftSquareBracket,
        RightSquareBracket,

        Make,
        Append,
        Semicolon,
        Package,
        Import,

        Comma,
        FullStop,

        In,
        Const,

        //line separators
        Tab,
        MultiLineStart,
        MultiLineEnd,
        Comment,
        EndComment,
        NewLine,

        NA, //invalid type
    }

    public enum ASTNodeType
    {
        Invalid,
        Comment,

        //Keywords
        Return,
        Break,
        Continue,
        Print,
        Println,

        //Literals
        IntegerLiteral,
        FloatLiteral,
        StringLiteral,
        CharLiteral,
        BoolLiteral,

        //__Operators__

        Minus,
        Reference,
        DereferenceAssignment,


        //__Bodies__
        FunctionBody,
        ForBody,
        ElseBody,
        IfBody,
        WhileBody,

        //__declarations__
        ArrayDeclaration,
        PointerDeclaration,
        FunctionDeclaration,
        Declaration,

        //__Functions__
        FunctionCall,

        //arrays
        ArrayGroup, //for multidimensional arrays
        ArrayElement,
        ArrayAccess,
        Array,


        //__Expressions__
        PrintExpression,
        BinaryExpression,
        ReturnExpression,
        BoolExpression,
        BoolComparison,

        //__Vartypes__
        VarType,
        Pointer,
        ReturnType,
        Const,
        Parameter,
        Parameters,

        //__Struct__
        StructVariable,


        //__Block things__
        IfStatement,
        WhileLoop,
        ForLoop,
        Else,

        //__General__
        Identifier,
        Assignment,

    } 

    public enum VarCategory
    {
        Regular,
        Array,
        Pointer,
    }

    public enum CharacterAction
    {
        Return,
        AddChar,
        Operator,
        AddCharAndBreak,
        Break,
        ReturnError,
    }

    public enum PotentialType
    {
        NA,
        String,
        Operator,
        Seperator,
        SeperatorToken, //brackets should be saved as tokens
        Number,
        Multiply,
        CharValue,
    }

    public enum LastTokenType
    {
        Null,
        Identifier,
        CustomVarType,
        Vartype,
        LeftSqBracket,
        RightSqBracket,
        LeftParenth,
        RightParenth,
        Comma,
        Semicolon,
        Newline,
        IntegerValue,
        Map,
        Pointer,
        FullStop,
        Interface,
        LeftBrace,
        RightBrace,
    }

    public enum DeclarationPhase
    {
        None,
        Comma,
        Semicolon,
        Type,
        Identifier,
    }
}
