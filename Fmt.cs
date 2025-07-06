

namespace TypeGo
{

    //It's quicker to Fmt.Println() than it is to write Console.WriteLine()
    public static class Fmt
    {
        public static void Print(string input)
        {
            Console.Write(input);
        }
        public static void Print(char input)
        {
            Console.Write(input);
        }
        public static void Println(string input)
        {
            Console.WriteLine(input);
        }
        public static void Println(char input)
        {
            Console.WriteLine(input);
        }
        public static void Println()
        {
            Console.WriteLine();
        }
        public static void PrintlnRed(string input)
        {
            Console.ForegroundColor = ConsoleColor.Red;
            Console.WriteLine(input);
            Console.ResetColor();
        }

        public static void PrintRed(string input)
        {
            Console.ForegroundColor = ConsoleColor.Red;
            Console.Write(input);
            Console.ResetColor();
        }
        public static void PrintlnYellow(string input)
        {
            Console.ForegroundColor = ConsoleColor.Yellow;
            Console.WriteLine(input);
            Console.ResetColor();
        }

        public static void PrintYellow(string input)
        {
            Console.ForegroundColor = ConsoleColor.Yellow;
            Console.Write(input);
            Console.ResetColor();
        }

        public static void PrintColor(string input, ConsoleColor color)
        {
            Console.ForegroundColor = color;
            Console.Write(input);
            Console.ResetColor();
        }
        public static void PrintlnColor(string input, ConsoleColor color)
        {
            Console.ForegroundColor = color;
            Console.WriteLine(input);
            Console.ResetColor();
        }
    }
}
