# CopyFile

Реализовать утилиту копирования файлов (см man dd). Выводить в консоль прогресс копирования. Программа
должна корректно обрабатывать ситуацию, когда offset или offset+limit за пределами source файла.
Пример использования:
# копирует 2К из source в dest, пропуская 1K данных
$ gocopy -from /path/to/source -to /path/to/dest -offset=1024 -limit=2048

go run main.go -from 12.txt -to out -offset=100 -limit=1000 
