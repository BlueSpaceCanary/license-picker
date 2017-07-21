#!/usr/bin/env python3.6

class Node:
    """Node in a tree."""
    #
    def __init__(self, text=None, path_text=None, children=[]):
        self.text = text
        self.path_text = path_text
        self.children = children

    def pick(self, choice):
        choices = list(filter(lambda x: x.path_text == choice, self.children))
        if choice == "help" or len(choices) == 0:
            print("Options:")
            for child in self.children:
                print(f"- {child.path_text}")

            return None


        else:
            return choices[0]


def setup():
    return Node(text="How to pick an open source license",
                children=[
                    Node(text="Do you care?",
                         path_text="Let's go",
                         children=[

                             Node(text="About non-US folks?",
                                  path_text="sure",
                                  children=[

                                      Node(text="Public Domain",
                                           path_text="HAHA NO",
                                           ),

                                      Node(text="Patents?",
                                           path_text="sure",
                                           children=[
                                               Node(text="Copyleft",
                                                    path_text="LOL",
                                                    children=[

                                                        Node(text="Bernie or Lenin?",
                                                             path_text="YES",
                                                             children=[

                                                                 Node(text="LGPL v2",
                                                                      path_text="Bernie",
                                                                      ),

                                                                 Node(text="GPL v2",
                                                                      path_text="Lenin",
                                                                      ),
                                                             ],
                                                             ),

                                                        Node(text="Can companies say you like their product",
                                                             path_text="NOPE",
                                                             children=[

                                                                 Node(text="BSD 3 Clause",
                                                                      path_text="What OMG no!",
                                                                      ),

                                                                 Node(text="MIT/X11/ZLIB",
                                                                      path_text="Why are you still asking me questions?",
                                                                      ),
                                                             ],
                                                             ),

                                                    ],
                                                    ),

                                               Node(text="Copyleft?",
                                                    path_text="Oh shit.",
                                                    children=[

                                                        Node(text="Apache 2.0",
                                                             path_text="NOPE",
                                                             ),

                                                        Node(text="Bernie or Lenin?",
                                                             path_text="YES",
                                                             children=[

                                                                 Node(text="LGPL v3",
                                                                      path_text="Bernie",
                                                                      ),

                                                                 Node(text="aGPL v3",
                                                                      path_text="Mao",
                                                                      ),

                                                                 Node(text="GPL v3",
                                                                      path_text="Lenin",
                                                                      ),
                                                             ],
                                                             ),

                                                    ],
                                                    ),

                                           ],
                                           ),
                                  ],
                                  ),

                             Node(text="WTFPL/CC0",
                                  path_text="NO",
                                  ),

                         ],
                         ),
                ],
                )


def main():
    """Do the thing."""
    node = setup()

    while node.children != []:
        choice = input(f"{node.text} (or 'help' to list options): ")
        result = node.pick(choice)
        if result is not None:
            node = result

    print(node.text)


# could use pytest but that'd be work
def test():
    """Test the thing."""
    node = setup()
    n1 = node.pick("Let's go")
    assert n1.text == "Do you care?"
    assert n1.pick("NO").text == "WTFPL/CC0"

    n2 = n1.pick("sure")
    assert n2.text == "About non-US folks?"
    assert n2.pick("HAHA NO").text == "Public Domain"

    n3 = n2.pick("sure")

    # ...

    assert n3.pick("LOL").pick("YES").pick("Bernie").text == "LGPL v2"
    assert n3.pick("Oh shit.").pick("YES").pick("Lenin").text == "GPL v3"

    print("passed")


if __name__ == "__main__":
    main()
