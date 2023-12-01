using System;
using System.Linq;
using System.Collections;
using System.Collections.Generic;

namespace net45
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
            ////Console.WriteLine(solution(new int[] { -1, -3 }));
            //Console.WriteLine(solution(new int[] {1,3,6,4,1,2 }));
            //Console.WriteLine(solution(new int[] {1,2,3 }));

            var parser = new Parser();

            var output = parser.Parse(@"
Value:    True ;
Password:""""super%^&*333password;
DNSName:SomeName;

TimeToLive:4;
ClusterSize:2;
PortNumber:-222;

IsEnabled:true;
EnsureTransaction:false;
PersistentStorage:false;
");
        }

        public static int solution(int[] A)
        {
            // write your code in C# 6.0 with .NET 4.5 (Mono)
            var min = 1;
            foreach (var number in A.OrderBy(x => x))
            {
                if (number == min)
                {
                    min++;
                    continue;
                }
                else if (number < min)
                {
                    continue;
                }
                break;
            }
            return min;
        }

    }
    public class UnknownKeyException : Exception { };
    public class EmptyKeyException : Exception { };
    public class InvalidKeyException : Exception { };
    public class DuplicatedKeyException : Exception { };
    public class Parser
    {
        /*
        Here is an example input configuration to parse:

        User Name:admin;
        Password:""super%^&*333password;
        DNSName:SomeName;

        TimeToLive:4;
        ClusterSize:2;
        PortNumber:-222;

        IsEnabled:true;
        EnsureTransaction:false;
        PersistentStorage:false;
        */
        private static System.Text.RegularExpressions.Regex configLineRegex = new System.Text.RegularExpressions.Regex(@"(.*){0,1}:(.*){0,1};");
        private static System.Text.RegularExpressions.Regex validKey = new System.Text.RegularExpressions.Regex(@"^[a-zA-Z0-9]+$");
        private static System.Text.RegularExpressions.Regex validBool = new System.Text.RegularExpressions.Regex(@"^(true|false)$", System.Text.RegularExpressions.RegexOptions.IgnoreCase);
        private static System.Text.RegularExpressions.Regex validInt = new System.Text.RegularExpressions.Regex(@"^\-{0,1}[0-9]+$");
        public dynamic Parse(string configuration)
        {
            if (string.IsNullOrEmpty(configuration))
                throw new ArgumentException();
            var dictionary = new Dictionary<string, object>();

            foreach (var line in configuration.Split("\r\n".ToCharArray(), StringSplitOptions.RemoveEmptyEntries))
            {
                var match = configLineRegex.Match(line);
                string key = string.Empty;
                string value = string.Empty;
                string[] split = line.Split(':');
                if(split.Length == 2)
                {
                    key = split[0].Trim();
                    value = split[1].TrimEnd(';');
                }
                else
                {
                    key = split[0].TrimEnd(';').Trim();
                }

                if (string.IsNullOrEmpty(key))
                {
                    throw new EmptyKeyException();
                }

                if (!validKey.IsMatch(key) || char.IsDigit(key[0]))
                {
                    throw new InvalidKeyException();
                }

                if (dictionary.ContainsKey(key))
                {
                    throw new DuplicatedKeyException();
                }

                if (validBool.IsMatch(value))
                {
                    dictionary.Add(key, bool.Parse(value));
                }
                else if (validInt.IsMatch(value))
                {
                    dictionary.Add(key, int.Parse(value));
                }
                else
                {
                    dictionary.Add(key, value.Trim());
                }
            }

            return new DynamicDictionary(dictionary);
        }

        public class DynamicDictionary : System.Dynamic.DynamicObject
        {
            private Dictionary<string, object> keyValues;
            public DynamicDictionary(Dictionary<string, object> dictionary)
            {
                this.keyValues = dictionary;
            }

            public override bool TryGetMember(System.Dynamic.GetMemberBinder binder, out object result)
            {
                if (!this.keyValues.ContainsKey(binder.Name))
                {
                    throw new UnknownKeyException();
                }
                result = this.keyValues[binder.Name];
                return true;
            }
        }
    }
}
