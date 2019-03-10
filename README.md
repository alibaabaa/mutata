# Mutata 

## Get an md5 to stdout

```
> mutata -input=test md5
098f6bcd4621d373cade4e832627b4f6
```

## Get a bcrypt to file

```
> mutata -input=test -outfile=file.dat bcrypt
> cat file.dat
$2y$04$IpATSc1115UFXBYnMnr5pOvMeE0cdrKXb1D0n9f7dGAVoQvspqnAe
```

## Get an sha1 of 32 bytes of random data to stdout

```
> mutata -input=rand(32) sha1
Random (hex)
74 a3 96 fd 2c 9b ac 95 ed 9b ff f5 2d 77 ed cf 
a3 9e 9d cd 7c 47 12 6e f5 4c 53 90 c6 4e 4b bc 
a7eac69eced8c610dc95abacc4406c899dfa2662
```

## Get an md5 in base64 to stdout

```
> mutata -input=test md5 base64
MDk4ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzMjYyN2I0ZjY=
```

## Get an md5 in base64 of 32 bytes of random data to stdout

```
> mutata -input=rand(32) md5 base64
Random (hex)
91 28 17 fd c4 e6 5b f7 3f 11 ba c5 74 28 cd c2 
81 98 e3 32 10 9f 37 c9 2f bb 9b dc 0d 22 ca dc 
OHNqZGhKSEE5OElTa2xhejc0c2dka14qenxraXNrYTE=
```

## Get md5 of file.txt to stdout
```
> mutata -infile=file.txt md5
098f6bcd4621d373cade4e832627b4f6
```

## Get three rounds of md5 of file.txt to output.txt
```
> mutata -infile=file.txt -outfile=output.txt md5 md5 md5
098f6bcd4621d373cade4e832627b4f6
```