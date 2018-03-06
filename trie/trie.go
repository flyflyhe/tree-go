package main

import (
	"strconv"
	"strings"
	"bufio"
	"os"
	"fmt"
)

type HashSet struct {
	Set map[int]bool
}

func NewHashSet() *HashSet {
	return &HashSet{make(map[int]bool)}
}

func (set *HashSet) Add(i int) bool {
	_, found := set.Set[i]
	set.Set[i] = true
	return !found //False if it existed already
}

func (set *HashSet) Get(i int) bool {
	_, found := set.Set[i]
	return found //true if it existed already
}

func (set *HashSet) Remove(i int) {
	delete(set.Set, i)
}

type TrieNode struct {

	//子节点
	ChildNodes map[int]*TrieNode

	//词频统计
	Freq int

	//记录该节点的字符
	NodeChar string

	//插入记录时的编号id
	HashSet *HashSet
}

func (t * TrieNode) init() {
	t.ChildNodes = make(map[int]*TrieNode)
}

func (t * TrieNode) AddTrieNode(word string, id int) {

	if word == "" {
		return
	}

	wordRune := []rune(word)
	
	//求字符地址，方便将该字符放入到26叉树中的哪一叉中
	k := int(wordRune[0]) - int('a')

	if _, ok := t.ChildNodes[k]; !ok  {
		t.ChildNodes[k] = new(TrieNode)
		t.ChildNodes[k].init()

		t.ChildNodes[k].HashSet = NewHashSet()
		t.ChildNodes[k].NodeChar = string(word[0])
	}

	//该id途径的节点
	t.ChildNodes[k].HashSet.Add(id)

	nextWord := wordRune[1:]

	if string(nextWord) == "" {
		t.ChildNodes[k].Freq += 1
	}

	t.ChildNodes[k].AddTrieNode(string(nextWord), id)
}

func (t * TrieNode) DeleteTrieNode(word string, id int) {
	if word == "" {
		return
	}

	wordRune := []rune(word)

	k := int(wordRune[0]) - int('a')

	if _, ok := t.ChildNodes[k]; !ok  {
		return
	}

	nextWord := wordRune[1:]

	if string(nextWord) == "" && t.ChildNodes[k].Freq > 0 {
		t.ChildNodes[k].Freq -= 1
	}

	t.ChildNodes[k].HashSet.Remove(id)

	t.DeleteTrieNode(string(nextWord), id)
}

// 检索单词的前缀,返回改前缀的Hash集合
func (t * TrieNode) SearchTrie(word string, hashSet *HashSet) *HashSet {
	if word == "" {
		return hashSet
	}

	wordRune := []rune(word)

	k := int(wordRune[0]) - int('a')


	nextWord := wordRune[1:]

	if string(nextWord) == "" {
		//采用动态规划的思想，word最后节点记录这途径的id

		if _, ok := t.ChildNodes[k]; !ok {
			return hashSet
		}

		for i, _ := range t.ChildNodes[k].HashSet.Set {
			hashSet.Add(i)
		}
	}

	if _, ok := t.ChildNodes[k]; !ok {
		return hashSet
	}

	t.ChildNodes[k].SearchTrie(string(nextWord), hashSet)

	return hashSet	
}

func main () {
	t := new(TrieNode)
	t.init()

	file, _ := os.Open("/tmp/1.txt")
	defer file.Close()
	
	inputReader := bufio.NewReader(file)
    for {
		inputString, err := inputReader.ReadString('\n')
		
		strArr := strings.Split(inputString, " ")

		if err != nil {
            break
		}     
	
		inputString = strings.TrimSpace(strArr[1])
		
		number, _ := strconv.Atoi(strArr[0])
		t.AddTrieNode(strings.ToLower(strArr[1]),  number)
	}

	hashSet := t.SearchTrie("go", NewHashSet())

	fmt.Println(hashSet)
	
	for i, _ := range hashSet.Set {
		fmt.Print("当前字符串的编号ID为")
		fmt.Println(i)
	}
}