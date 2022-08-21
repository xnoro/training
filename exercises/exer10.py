# desafio 1
# encontrar o índice de 'o' dentro da variável biblioteca
biblioteca = 'Biblioteca'
print(biblioteca.rindex('o'))

# desafio 2
# usando as palavras
'Desenvolvimento De Aplicações'
#imprima apenas a palavra 'De Aplicações'

frase = 'Desenvolvimento De Aplicações'
# print(frase.rindex('D'))
# print(frase[16:])
indice_d = frase.rindex('D')
indice_s = frase.rindex('s')
print(frase[indice_d:indice_s+1])

