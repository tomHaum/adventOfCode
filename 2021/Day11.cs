using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

using System.Drawing;
namespace AdventOfCode2021
{
    internal class Day11
    {

        internal class octopus
        {
            private static int X_MAX = 10;
            private static int Y_MAX = 10;
            internal octopus(int x, int y, char c)
            {
                this.X = x;
                this.Y = y;
                this.Light = int.Parse(c.ToString());
                this.ToFlash = false;
                this.Flashed = false;
            }
            public int X;
            public int Y;
            public bool ToFlash;
            public bool Flashed;

            public int Light;

            public bool Increment()
            {
                Light++;
                if (Light > 9)
                {
                    ToFlash = true;
                    Light = 9;
                }
                return ToFlash;
            }

            public IEnumerable<Point> Flash()
            {
                //checkLeft
                if (this.X != 0)
                    yield return new Point(X - 1, Y);
                //checkRight
                if (this.X != X_MAX - 1)
                    yield return new Point(X + 1, Y);
                //checkTop
                if (this.Y != 0)
                    yield return new Point(X, Y - 1);
                //checkBottom
                if (this.Y != Y_MAX - 1)
                    yield return new Point(X, Y + 1);

                //diagonals
                //topLeft
                if (this.X != 0 && this.Y != 0)
                    yield return new Point(X - 1, Y - 1);
                //bottomLeft
                if (this.X != 0 && this.Y != Y_MAX - 1)
                    yield return new Point(X - 1, Y + 1);
                if (this.X != X_MAX - 1 && this.Y != 0)
                    yield return new Point(X + 1, Y - 1);
                //bottomLeft
                if (this.X != X_MAX - 1 && this.Y != Y_MAX - 1)
                    yield return new Point(X + 1, Y + 1);

                this.Flashed = true;
            }

            public static void PrintSquids(octopus[,] octopi)
            {
                for (int x = 0; x < 10; x++)
                {
                    for (int y = 0; y < 10; y++)
                    {
                        Console.Write(octopi[x, y].Light);
                    }
                    Console.WriteLine();
                }
                Console.WriteLine();
            }

            public bool Reset()
            {
                if (Flashed)
                {
                    Flashed = false;
                    ToFlash = false;
                    Light = 0;
                    return true;
                }
                return false;
            }
        }

        internal static void Run()
        {
            var octopi = new octopus[10, 10];
            var lines = input.day11.Split("\r\n").ToArray();
            for (int x = 0; x < 10; x++)
            {
                for (int y = 0; y < 10; y++)
                {
                    octopi[x, y] = new octopus(x, y, lines[x][y]);
                }
            }
            int accum = 0;
            for (int i = 0; i < int.MaxValue; i++)
            {
                Queue<octopus> flashing = new Queue<octopus>();

                foreach (var o in octopi)
                {
                    if (o.Increment())
                        flashing.Enqueue(o);
                }
                var flashed = new HashSet<octopus>();
                while (flashing.Any())
                {
                    var o = flashing.Dequeue();
                    flashed.Add(o);
                    if (o.Flashed)
                        continue;

                    foreach (var n in o.Flash().Select(x => octopi[x.X, x.Y]))
                    {
                        if (!flashed.Contains(n))
                            if (n.Increment())
                                flashing.Enqueue(n);
                    }
                }

                int innerAccum = 0;

                foreach (var o in octopi)
                {
                    if (o.Reset())
                    {
                        innerAccum++;
                    }
                }
                accum += innerAccum;
                if (i == 99)
                {
                    Console.WriteLine($"Day 11 - Part 1: {accum}");
                }
                if(innerAccum == 100)
                {
                    Console.WriteLine($"Day 11 - Part 2: {i+1}");
                    break;
                }
            }
        }
    }
}
