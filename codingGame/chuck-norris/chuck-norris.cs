using System;
using System.Linq;
using System.IO;
using System.Text;
using System.Collections;
using System.Collections.Generic;

class Solution
{
    static void Main(string[] args)
    {
        string MESSAGE = Console.ReadLine();
        Console.Error.WriteLine(MESSAGE);
        string strBin = "";
        foreach (char c in MESSAGE)
        {
            strBin += Convert.ToString(c, 2).PadLeft(7, '0');
        }
        string buff = "";
        string buff2 = "";
        int count = 0;
        for (int i = 0; i < strBin.Length; i++)
        {
            if (buff.Length == 0)
            {
                buff += strBin[i];
            }
            if (strBin[i] == buff[buff.Length-1])
            {
                count++;
            }
            else
            {
                buff2 += buff == "1" ? "0 " : "00 " ;
                for (int k = 0; k < count; k++)
                {
                    buff2 += "0";
                }
                buff2 += " ";
                buff = "" + strBin[i];
                count = 1;
            }
        }
        buff2 +=  buff == "1" ? "0 " : "00 " ; 
        for (int k = 0; k < count; k++)
        {
            buff2 += "0";
        }
        Console.WriteLine(buff2);
    }
}
