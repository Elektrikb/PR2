# PR2

## Как использовать:
1) Скомпилируйте и запустите программу

2) Введите пути к файлам через пробел или оставьте поле пустым для анализа всех .txt файлов в текущей папке

3) Программа выведет результаты анализа для каждого файла и общие итоги

```
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;

namespace FileAnalyzer
{
    // Структура для хранения результатов анализа файла
    public struct FileAnalysis
    {
        public string FileName { get; set; }
        public int WordCount { get; set; }
        public int CharCount { get; set; }
    }

    class Program
    {
        private static readonly object _lock = new object();
        private static List<FileAnalysis> _results = new List<FileAnalysis>();

        static async Task Main(string[] args)
        {
            Console.WriteLine("Введите пути к файлам через пробел или оставьте пустым для анализа всех .txt файлов в папке:");
            string input = Console.ReadLine();

            string[] filePaths;
            if (string.IsNullOrWhiteSpace(input))
            {
                filePaths = Directory.GetFiles(Directory.GetCurrentDirectory(), "*.txt");
                if (filePaths.Length == 0)
                {
                    Console.WriteLine("Не найдено .txt файлов в текущей директории.");
                    return;
                }
            }
            else
            {
                filePaths = input.Split(new[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
            }

            Console.WriteLine("\nНачинаем анализ файлов...");

            var tasks = new List<Task>();
            foreach (var filePath in filePaths)
            {
                tasks.Add(AnalyzeFileAsync(filePath));
            }

            await Task.WhenAll(tasks);

            DisplayResults();
        }

        static async Task AnalyzeFileAsync(string filePath)
        {
            try
            {
                string content;
                using (var reader = File.OpenText(filePath))
                {
                    content = await reader.ReadToEndAsync();
                }

                var fileName = Path.GetFileName(filePath);
                int wordCount = content.Split(new[] { ' ', '\t', '\n', '\r' }, StringSplitOptions.RemoveEmptyEntries).Length;
                int charCount = content.Length;

                var analysis = new FileAnalysis
                {
                    FileName = fileName,
                    WordCount = wordCount,
                    CharCount = charCount
                };

                // Синхронизация доступа к общему списку результатов
                lock (_lock)
                {
                    _results.Add(analysis);
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine($"Ошибка при обработке файла {filePath}: {ex.Message}");
            }
        }

        static void DisplayResults()
        {
            Console.WriteLine("\nРезультаты анализа:");

            lock (_lock)
            {
                int totalWords = 0;
                int totalChars = 0;
                int fileNumber = 1;

                foreach (var result in _results.OrderBy(r => r.FileName))
                {
                    Console.WriteLine($"{fileNumber++}. {result.FileName}: {result.WordCount} слов, {result.CharCount} символов");
                    totalWords += result.WordCount;
                    totalChars += result.CharCount;
                }

                Console.WriteLine($"\nИтог: {totalWords} слов, {totalChars} символов.");
            }
        }
    }
}
```
