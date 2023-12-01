using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

using System.Drawing;
using System.Text.RegularExpressions;
namespace AdventOfCode2021
{
    internal class Day15
    {
        internal static void Run()
        {
            var graph = To2D<int>(input.day15.Split("\r\n").Select(x => x.Select(y => y - '0').ToArray()).ToArray());
            var xLen = graph.GetLength(1);
            var yLen = graph.GetLength(0);


            var dist = new int[yLen, xLen];
            var prev = new Point[yLen, xLen];

            var start = new Point(0, 0);
            var q = new List<Point>();

            for (int x = 0; x < xLen; x++)
            {
                for (int y = 0; y < yLen; y++)
                {
                    dist[y, x] = int.MaxValue;
                    prev[y, x] = new Point(-1, -1);
                    q.Add(new Point(x, y));
                }
            }

            dist[start.Y, start.X] = 0;

            while (q.Count > 0)
            {
                var u = q.OrderBy(u => dist[u.Y, u.X]).First();

                q.Remove(u);

                foreach (var v in u.GetNeighbors(xLen, yLen))
                {
                    if (!q.Contains(v))
                        continue;
                    var alt = dist[u.Y, u.X] + length(u, v, graph);
                    if (alt < dist[v.Y, v.X])
                    {
                        dist[v.Y, v.X] = alt;
                        prev[v.Y, v.X] = u;
                    }
                }
            }

            var path = GetPath(prev, start, new Point(xLen - 1, yLen - 1)).ToList();
            path.Reverse();

            var accum = 0;
            foreach (var p in path.Skip(1))
            {
                var cost = graph[p.Y, p.X];
                accum += cost;
                //Console.WriteLine($"Point [{p.X},{p.Y}] | Value [{cost}] | totalPathCost [{accum}]");
            }
            Console.WriteLine($"Day 15 - Part 1: {accum}");

            dist = new int[yLen * 5, xLen * 5];
            prev = new Point[yLen * 5, xLen * 5];

            start = new Point(0, 0);
            q = new List<Point>();

            for (int x = 0; x < xLen * 5; x++)
            {
                for (int y = 0; y < yLen * 5; y++)
                {
                    dist[y, x] = int.MaxValue;
                    prev[y, x] = new Point(-1, -1);
                    q.Add(new Point(x, y));
                }
            }

            dist[start.Y, start.X] = 0;

            while (q.Count > 0)
            {
                var u = q.OrderBy(u => dist[u.Y, u.X]).First();

                q.Remove(u);

                foreach (var v in u.GetNeighbors(xLen*5, yLen*5))
                {
                    if (!q.Contains(v))
                        continue;
                    var alt = dist[u.Y, u.X] + length2(v, graph, xLen, yLen);
                    if (alt < dist[v.Y, v.X])
                    {
                        dist[v.Y, v.X] = alt;
                        prev[v.Y, v.X] = u;
                    }
                }
            }

            path = GetPath(prev, start, new Point((xLen * 5) - 1, (yLen * 5) - 1)).ToList();
            path.Reverse();
            accum = 0;
            foreach (var p in path.Skip(1))
            {
                var cost = length2(p, graph, xLen, yLen);
                accum += cost;
                //Console.WriteLine("Point [{0,2},{1,2}] | Value [{2}] | totalPathCost [{3,5}]", p.X, p.Y, cost, accum);
            }

            Console.WriteLine($"Day 15 - Part 2: {accum}");
        }

        static IEnumerable<Point> GetPath(Point[,] prev, Point start, Point end)
        {
            yield return end;
            var previous = end;
            Point current = prev[end.Y, end.X];
            yield return current;
            while (current != start)
            {
                var temp = current;
                current = prev[current.Y, current.X];
                yield return current;
                previous = temp;
            }
        }

        static T[,] To2D<T>(T[][] source, bool flip = false)
        {
            try
            {
                int FirstDim = source.Length;
                int SecondDim = source.GroupBy(row => row.Length).Single().Key; // throws InvalidOperationException if source is not rectangular

                var result = new T[FirstDim, SecondDim];
                for (int i = 0; i < FirstDim; ++i)
                    for (int j = 0; j < SecondDim; ++j)
                        if (flip)
                            result[j, i] = source[i][j];
                        else
                            result[i, j] = source[i][j];

                return result;
            }
            catch (InvalidOperationException)
            {
                throw new InvalidOperationException("The given jagged array is not rectangular.");
            }
        }

        public static int length(Point u, Point v, int[,] values)
        {
            return values[v.Y, v.X];
        }

        public static int length2(Point v, int[,] values, int xMax, int yMax)
        {

            var val = values[v.Y % yMax, v.X % xMax];
            val += (v.Y / yMax) + (v.X / xMax);
            if (val > 9)
                val -= 9;
            return val;
        }
    }
    public static class Extensions
    {
        public static IEnumerable<Point> GetNeighbors(this Point point, int xMax, int yMax, bool includeDiagonals = false)
        {
            //checkLeft
            if (point.X != 0)
                yield return new Point(point.X - 1, point.Y);
            //checkRight
            if (point.X != xMax - 1)
                yield return new Point(point.X + 1, point.Y);
            //checkTop
            if (point.Y != 0)
                yield return new Point(point.X, point.Y - 1);
            //checkBottom
            if (point.Y != yMax - 1)
                yield return new Point(point.X, point.Y + 1);
            if (includeDiagonals)
            {

                //diagonals
                //topLeft
                if (point.X != 0 && point.Y != 0)
                    yield return new Point(point.X - 1, point.Y - 1);
                //bottomLeft
                if (point.X != 0 && point.Y != yMax - 1)
                    yield return new Point(point.X - 1, point.Y + 1);
                if (point.X != xMax - 1 && point.Y != 0)
                    yield return new Point(point.X + 1, point.Y - 1);
                //bottomLeft
                if (point.X != xMax - 1 && point.Y != yMax - 1)
                    yield return new Point(point.X + 1, point.Y + 1);
            }

        }

    }
}
