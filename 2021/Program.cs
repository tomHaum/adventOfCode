// See https://aka.ms/new-console-template for more information
using AdventOfCode2021;
using BenchmarkDotNet.Running;
using System.Diagnostics;

////use the second Core/Processor for the test
//Process.GetCurrentProcess().ProcessorAffinity = new IntPtr(2);

////prevent "Normal" Processes from interrupting Threads
//Process.GetCurrentProcess().PriorityClass = ProcessPriorityClass.High;

////prevent "Normal" Threads from interrupting this thread
//Thread.CurrentThread.Priority = ThreadPriority.Highest;

//Day01.Run();
//Day02.Run();
//Day03.Run();
//Day04.Run();
//Day05.Run();
//Day06.Run();
//Day07.Run();
//Day08.Run();
//Day09.Run();
//Day10.Run();
Day11.Run();
//var summary = BenchmarkRunner.Run<Day6>();