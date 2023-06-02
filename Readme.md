# Filestire CLI

A CLI interface to an HTTP file server with multiple. This CLI acts as a client which can add, remove, edit and list files on the server.

# Getting started

## Making the operator available on the in-cluster OperatorHub



## Usage

#### `add` command will upload the files mentioned in the arguments to the server. add command supports multiple file uploads at once

```
filestore add a.txt b.txt
```

#### `ls` command will list all the files present on the server

```
filestore ls
```

#### `rm` command will delete the file mentioned in the next argument. In the below example a.txt will be deleted

```
filestore rm a.txt
```

#### `update` command will update the file mentioned in the next argument. If the file's content are updated the same will be reflected on the server as well otherwise server's copy of the file will remain unchanged

```
filestore update a.txt
```
#### `wc` command will list the count of all the words present in all the files

```
filestore wc
```
#### `freq-words` command will list top 10 most frequently used words by default. You can change the order and number of responses

```
filestore freq-words [--limit|-n 10] [--order=dsc|asc]
```
