'''
一摞有序的纸牌
'''
import collections
Card = collections.namedtuple('Card',['rank','suit'])
class FrencDeck:
    ranks = [str(n) for n in range(2,11)] + list('JKQA')
    suits = 'spades diamonds clubs hearts'.split()
    def __init__(self):
        self._cards=[Card(rank,suit) for suit in self.suits for rank in self.ranks]

    def __len__(self):#len函数会自动调用该方法
        return len(self._cards)

    def __getitem__(self,position):  #该方法用以支持下标索引取值
        return self._cards[position]


desk=FrencDeck()

print(desk[10])
print(desk[0])
print(desk[12::14])
print(desk[-1])
for i in desk:
    print(i)
print(len(desk))
print(Card("Q","hearts") in desk)

suit_values = dict(spades=3,hearts=2, diamonds=1,clubs=0)

def spades_high(card):
    rank_value = FrencDeck.ranks.index(card.rank)
    return rank_value * len(suit_values) + suit_values[card.suit]
print(spades_high(Card("Q","hearts")),end='\n\n')

for card in sorted(desk,key=spades_high):
    print(card)




