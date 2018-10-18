using System;
using System.IO;

namespace protobufDemo
{
	class Program
	{
		static void Main(string[] args)
		{
			string rootFolderName = "protobufDemo";
			var ind = Directory.GetCurrentDirectory().ToString().IndexOf(rootFolderName);
			string rootFolder = Directory.GetCurrentDirectory().ToString().Substring(0, ind + rootFolderName.Length);

			// Read all protobuf encoded files
			foreach (var f in Directory.GetFiles(rootFolder, "*.dat"))
			{
				Contacts.AddressBook book;
				using (var input = File.OpenRead(f))
				{
					book = Contacts.AddressBook.Parser.ParseFrom(input);
				}

				Console.WriteLine("\n{0} - {1}", f, book);

				foreach (var p in book.People)
				{
					Console.WriteLine("\n{0} - {1}, {2}, {3}", f, p.Name, p.Email, p.Address);
				}
			}

			Console.ReadKey();
		}
	}
}
