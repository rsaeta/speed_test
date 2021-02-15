import UIKit

var str = "Hello, playground"

class Card {
    private let suit: Int
    private let rank: Int
    
    init(suit: Int, rank: Int) {
        self.suit = suit
        self.rank = rank
    }
    
    func sharesRank(other: Card) -> Bool {
        return other.rank == self.rank
    }
    
    func sharesSuit(other: Card) -> Bool {
        return other.suit == self.suit
    }
}

class Deck {
    private var deck: [Card]
    
    init() {
        self.deck = [Card]()
        for suit in 0...4 {
            for rank in 0...13 {
                self.deck.append(Card(suit: suit, rank: rank))
            }
        }
    }
    
    func shuffle() {
        self.deck.shuffle()
    }
    
    func getNext() -> Card? {
        return self.deck.popLast()
    }
}

class Hand {
    private var hand: [Card]
    
    init() {
        self.hand = [Card]()
    }
    
    func addCard(card: Card) {
        self.hand.append(card)
        while self.hand.count > 4 && self._applyRule() {
        }
    }
    
    private func _applyRule() -> Bool {
        let lastIndex = self.hand.count - 1
        let fourthIndex = self.hand.count - 4
        let topCard = self.hand[lastIndex]
        let compareTo = self.hand[fourthIndex]
        
        if topCard.sharesSuit(other: compareTo) {
            for _ in 0...4 {
                self.hand.remove(at: fourthIndex - 1)
            }
            return true
        }
        
        if topCard.sharesRank(other: compareTo) {
            self.hand.remove(at: fourthIndex + 1)
            self.hand.remove(at: fourthIndex + 1)
            return true
        }
        return false
    }
    
    func size() -> Int {
        return hand.count
    }
}

/**
 * Plays an iteration of game and returns
 * number left on the stack.
 */
func playIteration() -> Int {
    let deck = Deck()
    let hand = Hand()
    
    deck.shuffle()
    var card = deck.getNext()
    while card != nil {
        hand.addCard(card: card!)
        card = deck.getNext()
    }
    
    return hand.size()
}

var results = [Int]()
for _ in 1...10000 {
    results.append(playIteration())
}

results.max()


