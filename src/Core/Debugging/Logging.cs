

using System.Text;

namespace TypeGo
{
    public static class Logging
    {
        public static bool DEBUG = true;

        public static StringBuilder Log = new StringBuilder();

        public static void CreateLog() {
            Log = new StringBuilder();
        }

        public static string GetLog() {
            return Log.ToString();
        }

        public static void ClearLog() {
            Log.Clear();
        }

        public static void PrintLog() {
            if (DEBUG == false) {
                return;
            }
            Fmt.Println("\nPrinting log:");
            Fmt.Println(GetLog());
        }

        public static void AddToLog(string message)
        {
            if (DEBUG == false) {
                return;
            }
            Log.AppendLine(message);
        }
    }
}
