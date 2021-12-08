using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace AdventOfCode2021
{
    internal class Day4
    {
        private static Regex numbers = new Regex(@"(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)");
        internal static void Run()
        {
            var txt = input.day4.Split("\r\n");

            var randomOrder = txt[0].Split(',').Select(x => int.Parse(x)).ToArray();

            var boards = new List<int[][]>();
            var marked = new List<bool[][]>();

            var currBoard = new int[5][];
            var currMarked = new bool[5][];
            var i = 0;
            foreach (var line in txt.Skip(2))
            {
                if(line == "")
                {
                    boards.Add(currBoard);
                    marked.Add(currMarked);
                    currBoard = new int[5][];
                    currMarked = new bool[5][];
                    i = 0;
                    continue;
                }

                var matches = numbers.Match(line);
                currMarked[i] = Enumerable.Range(0, 5).Select(x => false).ToArray();
                currBoard[i++] = matches.Groups.Values.Skip(1).Select(x => int.Parse(x.Value)).ToArray();
            }
            boards.Add(currBoard);
            marked.Add(currMarked);

            var winners = new HashSet<int>();
            var winner = false;
            foreach(var num in randomOrder)
            {
                //loop boards
                for(int b =0; b < boards.Count; b++)
                {
                    if (winners.Contains(b))
                        continue;
                    var board = boards[b];
                    var mark = marked[b];
                    //loop x
                    for(int x = 0; x < 5; x++)
                    {
                        //loop y
                        for(int y = 0; y < 5; y++)
                        {
                            if(board[x][y] == num)
                            {
                                mark[x][y] = true;
                            }
                            int accum = 0;
                            // check winner
                            
                            //  horizontal
                            accum = 0;
                            for(i = 0; i < 5; i++)
                            {
                                if (mark[x][i])
                                    accum++;
                            }
                            if(accum == 5)
                            {
                                // winner
                                //Console.WriteLine("WINNER");
                                winner = true;
                                break;
                            }
                            accum = 0;

                            //  vertical
                            accum = 0;
                            for (i = 0; i < 5; i++)
                            {
                                if (mark[i][y])
                                    accum++;
                            }
                            if (accum == 5)
                            {
                                // winner
                                //Console.WriteLine("WINNER");
                                winner = true;
                                break;
                            }
                            ////  left diag
                            //accum = 0;
                            //for (i = 0; i < 5; i++)
                            //{
                            //    if (mark[i][i])
                            //        accum++;
                            //}
                            //if (accum == 5)
                            //{
                            //    // winner
                            //    Console.WriteLine("WINNER");
                            //}
                            ////  right diag
                            //accum = 0;
                            //for (i = 0; i < 5; i++)
                            //{
                            //    if (mark[5-i-1][i])
                            //        accum++;
                            //}
                            //if (accum == 5)
                            //{
                            //    // winner
                            //    Console.WriteLine("WINNER");
                            //}
                        }
                    }

                    if(winner)
                    {

                        if (winners.Count == 0 || winners.Count == boards.Count -1)
                        {

                            var sum = 0;
                            for (int x = 0; x < 5; x++)
                            {
                                for (int y = 0; y < 5; y++)
                                {
                                    if (!mark[x][y])
                                    {
                                        sum += board[x][y];
                                    }
                                }
                            }
                            Console.WriteLine($"Day 4 - Part {((winners.Count == 0) ? '1' : '2')} : Sum [{sum}] | last Num [{num}] | product [{num * sum}]");
                        }
                        winners.Add(b);
                        winner = false;
                    }
                }
            }
        }
    }
}
