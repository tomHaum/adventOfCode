using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
namespace AdventOfCodeCSharp
{
    class Day4
    {
        private static char[] sepertors = (Environment.NewLine + ' ').ToCharArray();
        private static Regex height = new Regex(@"(\d+)(in|cm)");
        private static Regex hair = new Regex(@"#[0-9a-f]{6}");
        private static Regex eye = new Regex(@"(amb|blu|brn|gry|grn|hzl|oth)");
        private static Regex passport = new Regex(@"^[0-9]{9}$");
        public static void Run()
        {
            Console.WriteLine("Day 4");
            var batches = data.day4.Split(Environment.NewLine + Environment.NewLine);

            int count = 0;

            foreach (var b in batches)
            {
                var tags = b.Split(sepertors, StringSplitOptions.RemoveEmptyEntries).Select(x => x.Substring(0, 3)).ToArray();

                if (tags.Length == 8 || tags.Length == 7 && !tags.Contains("cid"))
                {
                    count++;
                    continue;
                }
            }

            Console.WriteLine("Part 1: " + count);
            count = 0;

            foreach (var b in batches)
            {
                var tags = b.Split(sepertors, StringSplitOptions.RemoveEmptyEntries).Select(x => (tag: x.Substring(0, 3), value: x.Substring(4))).ToArray();
                //Console.WriteLine(tags[0].tag + " : " + tags[0].value);
                bool hasCID = false;
                bool valid = true;

                foreach (var t in tags)
                {
                    switch (t.tag)
                    {
                        case "byr":
                            var byr = int.Parse(t.value);
                            if (byr < 1920 || byr > 2002) 
                                valid = false;
                            break;
                        case "iyr":
                            var iyr = int.Parse(t.value);
                            if (iyr < 2010 || iyr > 2020) 
                                valid = false;
                            break;
                        case "eyr":
                            var eyr = int.Parse(t.value);
                            if (eyr < 2020 || eyr > 2030) 
                                valid = false;
                            break;
                        case "hgt":
                            var m = height.Match(t.value);
                            if (!m.Success)
                            {
                                valid = false;
                                break;
                            }
                            var cmOrIn = m.Groups[2].Value;
                            var hgt = int.Parse(m.Groups[1].Value);

                            switch (cmOrIn)
                            {
                                case "cm":
                                    if (hgt < 150 || hgt > 193)
                                        valid = false;
                                    break;
                                case "in":
                                    if (hgt < 59 || hgt > 76)
                                        valid = false;
                                    break;
                            }
                            break;
                        case "hcl":
                            var hcl = hair.Match(t.value);
                            if (!hcl.Success) 
                                valid = false;
                            break;
                        case "ecl":
                            var ecl = eye.Match(t.value);
                            if (!ecl.Success) 
                                valid = false;
                            break;
                        case "pid":
                            var pid = passport.Match(t.value);
                            if (!pid.Success) 
                                valid = false;
                            break;
                        case "cid":
                            hasCID = true;
                            break;
                        default:
                            valid = false;
                            break;
                    }
                    if (!valid)
                        break;
                }
                if (tags.Length < 7)
                    valid = false;
                if (tags.Length > 8)
                    valid = false;
                if (tags.Length == 7 && hasCID)
                    valid = false;
                if (valid) count++;
            }

            Console.WriteLine("Part 2: " + count);
        }
    }
}
