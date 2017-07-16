package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type node interface {
	ChoiceText() string
	NodeText() string
}

type fork struct {
	choice_text string
	node_text   string
	children    []node
}

func (f *fork) ChoiceText() string {
	return f.choice_text
}

func (f *fork) NodeText() string {
	return f.node_text
}

type leaf struct {
	node_text   string
	choice_text string
}

func (l *leaf) ChoiceText() string {
	return l.choice_text
}

func (l *leaf) NodeText() string {
	return l.node_text
}

func (f *fork) addChild(n node) {
	f.children = append(f.children, n)
}

// #go
func (f *fork) getMatch(in string) (interface{}, error) {
	for _, n := range f.children {
		if fmt.Sprintf("%s\n", n.ChoiceText()) == in {
			return n, nil
		}
	}

	return nil, errors.New("No match")
}

func (f *fork) walk() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s? (or \"help\" to list options): ", f.node_text)
	txt, _ := reader.ReadString('\n')

	m, err := f.getMatch(txt)
	if err != nil {
		fmt.Println("Options: ")
		for _, n := range f.children {
			fmt.Println(n.ChoiceText())
		}

		f.walk()
		return
	}

	switch t := m.(type) {
	default:
		panic("go away you jerk")
	case (*fork):
		t.walk()
	case (*leaf):
		fmt.Printf("*****%s*****n", t.node_text)
	}
}

func lolsetup() fork {
	// setup lol
	co_bs := fork{node_text: "Can companies say you like their product?", choice_text: "NOPE"}
	co_bs.addChild(&leaf{"BSD 3 clause", "What OMG no!"})
	co_bs.addChild(&leaf{"MIT/X11/zlib", "Why are you still asking me questions?"})

	v2 := fork{node_text: "Bernie or Lenin", choice_text: "YES"}
	v2.addChild(&leaf{"LGPLv2", "Bernie"})
	v2.addChild(&leaf{"GPLv2", "Lenin"})

	// same
	copyleft_ip_anarchy := fork{node_text: "Copyleft", choice_text: "LOL"}
	copyleft_ip_anarchy.addChild(&co_bs)
	copyleft_ip_anarchy.addChild(&v2)

	v3 := fork{node_text: "Bernie or Lenin", choice_text: "YES"}
	v3.addChild(&leaf{"LGPLv3", "Bernie"})
	v3.addChild(&leaf{"GPLv3", "Lenin"})
	v3.addChild(&leaf{"aGPLv3", "Mao"})

	apache := leaf{"Apache 2.0", "NOPE"}

	copyleft_ip_tyrrany := fork{node_text: "Copyleft", choice_text: "oh shit"}
	copyleft_ip_tyrrany.addChild(&apache)
	copyleft_ip_tyrrany.addChild(&v3)

	patents := fork{node_text: "patents?", choice_text: "sure"}
	patents.addChild(&copyleft_ip_anarchy)
	patents.addChild(&copyleft_ip_tyrrany)

	pd := leaf{"Public domain", "HAHA NO"}

	us_is_god := fork{node_text: "About non-US folks?", choice_text: "sure"}
	us_is_god.addChild(&pd)
	us_is_god.addChild(&patents)

	the_good_ones := leaf{"WTFPL/CC0", "NO"}

	meirl := fork{node_text: "Do you care?", choice_text: "Let's go"}

	meirl.addChild(&the_good_ones)
	meirl.addChild(&us_is_god)

	root := fork{node_text: "How to pick an open source license"}
	root.addChild(&meirl)

	return root
}

func main() {
	root := lolsetup()
	root.walk()
}
