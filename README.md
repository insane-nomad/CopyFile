# CopyFile

����������� ������� ����������� ������ (�� man dd). �������� � ������� �������� �����������. ���������
������ ��������� ������������ ��������, ����� offset ��� offset+limit �� ��������� source �����.
������ �������������:
# �������� 2� �� source � dest, ��������� 1K ������
$ gocopy -from /path/to/source -to /path/to/dest -offset=1024 -limit=2048

go run main.go -from 12.txt -to out -offset=100 -limit=1000 
