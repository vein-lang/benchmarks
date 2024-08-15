#!/usr/bin/env dotnet-script
var strValue = Args[0];
if (!int.TryParse(Args[1], out int size))
{
    Console.WriteLine("Invalid size: " + Args[1]);
    return;
}

for (int i = 0; i < size; i++)
{
    Console.WriteLine(strValue);
}