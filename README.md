# Часть 1

1. [Референсный геном e.coli](/part1/ref_e_coli.fna.gz)
2. [Результат сиквенирования e.coli](/part1/e_coli_fastq.gz)
3. [Ссылка на загруженные прочтения из NCBI SRA](https://www.ncbi.nlm.nih.gov/sra/SRX17981169[accn])
4. [Bash скрипт с реализованным алгоритмом](/part1/start.sh)
5. [Результат команды samtools ﬂagstat](/part1/result.txt)
6. [Скрипт разбора файлов с результатами](/part1/parser.py)

# Часть 2

## Инструкция по развертыванию и установке фреймворка:

1. Проверить, что установлен Go 1.17 или выше. [Инструкция по установке Go](https://tecadmin.net/install-go-on-ubuntu/)

```
go version
```

_Ответ должен быть `go version go1.17.7 linux/amd64`_

2. Установить библиотеку scipipe Go (не пропустите точки в конце!).

```
go install github.com/scipipe/scipipe/...@latest
```

## Запуск тестового пайплайна "Hello world":

1. Создать Go модуль для будущего скрипта:

```
go mod init myfirstworkflow-module
```

2. Сгенерировать hello_world пример с помощью команды:

```
scipipe new hello_world.go
```

_У меня не сработало, поэтому я вручную написал скрипт
[hello_world.go](/part2/hello_world.go)_

3. Убеждаемся, что все небходимые зависимости SciPipe установлены в модуле

```
go mod tidy
```

4. Запускаем скрипт

```
go run hello_world.go
```

## Результаты работы тестового пайплайна "Hello world":

1. Посмотреть на то, что сгенерировал SciPipe

```
ls -1 hello*
```

2. Для вывода результата используем

```
cat hello.out.txt.world.out.txt
```

В результате получаем [файл](/part2/hello.out.txt) и [фаил](/part2/hello.out.txt.world.out.txt), а также сопровождающий файл логов .audit.json для каждого из этих файлов. [например](/part2/hello.out.txt.world.out.txt.audit.json)

# Часть 3

1. В директории part3/ выполнить 'go run pipeline.go'
2. [Код](/part3/pipline.go) пайплайна на SciPipe
3. Все сгенерированные файлы в ходе работы пайплайна по умолчанию кладутся в папку output [пример](/part3/output/)
4. В результате в [файл](/part3/output/parsed.txt)выводится результат "OK" или "Not OK"
5. Лог файлы каждого запуска хранятся в [логи](/part3/logs/)
6. Визуализация пайплайна в виде графического файла:
