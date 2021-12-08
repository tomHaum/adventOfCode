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

//Day1.Run();
//Day2.Run();
//Day3.Run();
//Day4.Run();
//Day5.Run();
//Day6.Run();
//Day7.Run();
Day8.Run();
//var summary = BenchmarkRunner.Run<Day6>();