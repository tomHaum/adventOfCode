using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace AdventOfCodeCSharp
{
    class Day7
    {
        class Bag
        {

            public static Dictionary<string,Bag> AllBags = new Dictionary<string, Bag>();

            public List<(Bag, int)> ContainedBags = new List<(Bag, int)>();
            public List<(Bag, int)> BagsImIn = new List<(Bag, int)>();

            public string Name;
            public static void AddBag(string s)
            {
                var name = GetBagName(s);
                var bag = GetOrCreateBag(name);
                bag.ContainedBags = GetContains(s);

                foreach(var b in bag.ContainedBags)
                {
                    b.Item1.BagsImIn.Add((bag, b.Item2));
                }
            }
            private Bag(string name)
            {
                if (name.Contains("bag"))
                    Console.WriteLine("why");
                Name = name;
            }

            private static Regex MyNameReg = new Regex(@"^(\w+\s\w+)\sbags?");
            private static string GetBagName(string s)
            {
                var match = MyNameReg.Match(s);
                if (match.Success)
                    return match.Groups[1].Value;
                throw new ArgumentException("No name to this bag");
            }

            private static Regex ContainBagsReg = new Regex(@"(\d+)\s(\w+\s\w+)\sbags?");
            private static List<(Bag,int)> GetContains(string s)
            {
                var matches = ContainBagsReg.Matches(s);

                List<(Bag, int)> ret = new List<(Bag, int)>();

                foreach(Match x in matches)
                {
                    var bagName = x.Groups[2].Value;
                    var count = int.Parse(x.Groups[1].Value);
                    ret.Add(new(GetOrCreateBag(bagName), count));
                }

                return ret;
            }

            private static Bag GetOrCreateBag(string name)
            {
                if (AllBags.ContainsKey(name))
                    return AllBags[name];
                var bag = new Bag(name);
                AllBags.Add(name, bag);
                return bag;
            }

            public int GetBagCount()
            {
                var accum = 0;
                foreach(var bag in this.ContainedBags)
                {
                    accum += bag.Item2;
                    accum += bag.Item1.GetBagCount() * bag.Item2;
                }
                return accum;
            }
            public override string ToString()
            {
                var sb = new StringBuilder();
                sb.Append(this.Name);
                sb.Append(" contains ");
                
                for(int i = 0; i < this.ContainedBags.Count; i++)
                {
                    sb.Append(this.ContainedBags[i].Item2);
                    sb.Append(" ");
                    sb.Append(this.ContainedBags[i].Item1.Name);
                    sb.Append(" ");
                    if (this.ContainedBags[i].Item2 > 1)
                        sb.Append("bags");
                    else
                        sb.Append("bag");
                    if (i + 1 < this.ContainedBags.Count)
                        sb.Append(", ");
                }
                if (this.ContainedBags.Count == 0)
                    sb.Append("no bags");
                
                sb.Append(".");
                return sb.ToString();
            }
        }
        public static void Run()
        {
            Console.WriteLine("Day 7");

            var bags = data.day7.Split("\r\n");

            var bag1 = bags[0];

            //var bagName = Bag.GetBagName(bag1);
            //var subBags = Bag.GetContains(bag1);
            //Console.WriteLine(bagName);

            foreach (var bag in bags)
            {
                Bag.AddBag(bag);
            }
            var goldBag = Bag.AllBags["shiny gold"];

            var parentBags = new Queue<Bag>(goldBag.BagsImIn.Select(x => x.Item1));
            var ancestors = new HashSet<string>();
            while(parentBags.Count != 0)
            {
                var bag = parentBags.Dequeue();
                ancestors.Add(bag.Name);
                foreach(var b in bag.BagsImIn)
                {
                    parentBags.Enqueue(b.Item1);
                }
            }

            Console.WriteLine("Part 1: " + ancestors.Count);

            Console.WriteLine("Part 2: " + goldBag.GetBagCount());
            
        }
    }
}
