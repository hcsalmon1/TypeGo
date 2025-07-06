using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace TypeGo
{
    public static class VarUtils
    {
        public static bool IsVarType(Token token)
        {
            switch (token.Type)
            {
                case TokenType.Int:
                case TokenType.Int8:
                case TokenType.Int16:
                case TokenType.Int32:
                case TokenType.Int64:
                case TokenType.Uint:
                case TokenType.Uint8:
                case TokenType.Uint16:
                case TokenType.Uint32:
                case TokenType.Uint64:
                case TokenType.Float32:
                case TokenType.Float64:
                case TokenType.String:
                case TokenType.Byte:
                case TokenType.Rune:
                case TokenType.Bool:
                    return true;                    
                default:
                    return false;
            }
        }
        public static bool IsVarType(TokenType tokenType)
        {
            switch (tokenType)
            {
                case TokenType.Int:
                case TokenType.Int8:
                case TokenType.Int16:
                case TokenType.Int32:
                case TokenType.Int64:
                case TokenType.Uint:
                case TokenType.Uint8:
                case TokenType.Uint16:
                case TokenType.Uint32:
                case TokenType.Uint64:
                case TokenType.Float32:
                case TokenType.Float64:
                case TokenType.String:
                case TokenType.Byte:
                case TokenType.Rune:
                case TokenType.Bool:
                    return true;
                default:
                    return false;
            }
        }
    }
}
