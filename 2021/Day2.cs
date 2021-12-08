using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCode2021
{
    internal class Day2
    {
        internal class SubCommand
        {
            internal enum SubMovementDirection
            {
                Forward, Down, Up
            }

            public SubMovementDirection Direction;
            public int Amount;
        }
        internal static void Run()
        {
            var subCommands = input.day2.Split("\r\n").Select(x =>
            {
                var parts = x.Split(" ");

                int amount = int.Parse(parts[1]);
                SubCommand.SubMovementDirection direction;

                switch (parts[0])
                {
                    case "forward":
                        direction = SubCommand.SubMovementDirection.Forward;
                        break;
                    case "down":
                        direction = SubCommand.SubMovementDirection.Down;
                        break;
                    case "up":
                        direction = SubCommand.SubMovementDirection.Up;
                        break;
                    default:
                        throw new Exception($"Could not parse {parts[0]} into a SubCommand.Direction");
                }

                return new SubCommand {
                    Amount = amount,
                    Direction = direction
                };
            }).ToList();

            int horizontalPosition = 0;
            int depth = 0;

            foreach(var command in subCommands)
            {
                switch (command.Direction)
                {
                    case SubCommand.SubMovementDirection.Forward:
                        horizontalPosition += command.Amount;
                        break;
                    case SubCommand.SubMovementDirection.Down:
                        depth += command.Amount;
                        break;
                    case SubCommand.SubMovementDirection.Up:
                        depth -= command.Amount;
                        break;
                }
            }

            Console.WriteLine($"Day 2 - Part 1: hortizontalPosition [{horizontalPosition}] | depth[{depth}] | product [{horizontalPosition * depth}]");
            horizontalPosition = 0;
            depth = 0;
            int aim = 0;
            foreach (var command in subCommands)
            {
                switch (command.Direction)
                {
                    case SubCommand.SubMovementDirection.Forward:
                        horizontalPosition += command.Amount;
                        depth += aim * command.Amount;
                        break;
                    case SubCommand.SubMovementDirection.Down:
                        aim   += command.Amount;
                        break;
                    case SubCommand.SubMovementDirection.Up:
                        aim   -= command.Amount;
                        break;
                }
            }
            Console.WriteLine($"Day 2 - Part 1: hortizontalPosition [{horizontalPosition}] | depth[{depth}] | aim[{aim}]|product [{horizontalPosition * depth}]");
        }
    }
}
