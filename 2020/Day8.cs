using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCodeCSharp
{
    class Day8
    {
        public static void Run()
        {
            Console.WriteLine("Day 8");
            var program = data.day8.Split("\r\n");
            (var x, var y) = FindLoop(program);
            Console.WriteLine("Part 1: " + accum);

        }




        private static (int index,int accum) FindLoop(string[] program)
    {
        var visited = new bool[program.Length];

        var pc = 0;
        var accum = 0;
        while (!visited[pc] || pc > program.Length)
        {
            visited[pc] = true;
            Console.WriteLine(program[pc]);
            var cmd = program[pc].Substring(0, 3);
            var cntstr = program[pc].Substring(4);
            var cnt = int.Parse(cntstr);

            switch (cmd)
            {
                case "nop":
                    pc++;
                    break;
                case "jmp":
                    pc += cnt;
                    break;
                case "acc":
                    pc++;
                    accum += cnt;
                    break;
            }
        }
    }   
}
