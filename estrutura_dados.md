# Estrutura de Dados

Este documento descreve várias estruturas de dados e algoritmos comuns usados em programação e resolução de problemas.

## Estruturas de Dados

### Array
Um **array** é uma estrutura de dados que contém um número fixo de elementos do mesmo tipo, armazenados em blocos contínuos de memória. Acesso a qualquer elemento pode ser feito em tempo constante, `O(1)`, usando o índice.

### LinkedList
A **LinkedList** é uma coleção de elementos chamados nós. Cada nó contém dois campos: um valor armazenado e uma referência (ou ponteiro) para o próximo nó. Há duas variações principais:
- **Singly Linked List**: Cada nó aponta para o próximo nó.
- **Doubly Linked List**: Cada nó aponta para o próximo e o anterior.

### Queue
Uma **fila (Queue)** é uma estrutura de dados do tipo FIFO (First In, First Out), onde o primeiro elemento inserido é o primeiro a ser removido. As operações principais são:
- `enqueue`: adicionar um elemento ao final da fila.
- `dequeue`: remover um elemento da frente da fila.

### Stack
Uma **pilha (Stack)** é uma estrutura de dados do tipo LIFO (Last In, First Out), onde o último elemento inserido é o primeiro a ser removido. As operações principais são:
- `push`: adicionar um elemento ao topo da pilha.
- `pop`: remover o elemento do topo da pilha.

### Binary Tree
Uma **árvore binária** é uma estrutura de dados em que cada nó possui no máximo dois filhos, chamados de **filho esquerdo** e **filho direito**. Existem diferentes tipos de árvores binárias, como árvores de busca binária (BST), onde a ordem dos nós é tal que os valores menores ficam à esquerda e os maiores à direita.

### Hashmap
Um **Hashmap** é uma estrutura de dados que mapeia chaves para valores. Ele usa uma função de hash para calcular o índice da chave, permitindo busca, inserção e remoção eficientes em tempo médio `O(1)`.

### Graph
Um **grafo** é uma estrutura de dados composta por **nós (vértices)** e **arestas (conexões)**. Os grafos podem ser:
- **Dirigidos** ou **não dirigidos**.
- **Ponderados** ou **não ponderados**.

## Identificação de Tipos de Problemas

### DFS (Busca em Profundidade)
O **Depth First Search (DFS)** é um algoritmo de busca que explora o grafo começando por um nó e explora o mais profundo possível ao longo de cada ramo antes de retroceder.

### BFS (Busca em Largura)
O **Breadth First Search (BFS)** é um algoritmo de busca que explora todos os nós de um nível antes de avançar para o próximo nível. Muito usado em problemas de menor caminho.

### Sliding Window
A técnica **Sliding Window** é usada em problemas que envolvem subarrays ou subsequências contínuas, onde uma janela de tamanho variável é movida ao longo do array para encontrar a solução.

### Backtracking
O **Backtracking** é uma técnica de tentativa e erro que envolve explorar todas as soluções possíveis, recuando quando uma solução parcial se mostra inviável.

### DP (Programação Dinâmica)
**Programação Dinâmica (DP)** é uma técnica de otimização que resolve problemas complexos dividindo-os em subproblemas menores e armazenando seus resultados. Um exemplo clássico é o cálculo do **Fibonacci**.

### Hashmap
Usado para resolver problemas que requerem acesso rápido a dados, como verificar a existência de elementos em um conjunto, contagem de frequências, etc.

### Two Pointer
A técnica **Two Pointer** envolve o uso de dois ponteiros para explorar as soluções possíveis em um array ou lista. Usado para encontrar pares que atendam a uma condição específica.
