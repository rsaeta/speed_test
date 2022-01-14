// ✗ time node main.js
// Mean: 18.568024
// Stdev: 7.326537295038032
// node main.js  5.69s user 0.18s system 98% cpu 5.935 total

type card = number
type deck = card[]

function shuffle(d: deck): deck {
    var currentIndex = d.length, temporaryValue, randomIndex;

    // While there remain elements to shuffle...
    while (0 !== currentIndex) {
  
      // Pick a remaining element...
      randomIndex = Math.floor(Math.random() * currentIndex);
      currentIndex -= 1;
  
      // And swap it with the current element.
      temporaryValue = d[currentIndex];
      d[currentIndex] = d[randomIndex];
      d[randomIndex] = temporaryValue;
    }
  
    return d;
}

function playIteration(): number {
    let deck = shuffle(getDeck())
    let hand = []
    deck.forEach(card => {
        hand.push(card)
        hand = tryReduce(hand)
    })

    return hand.length
}

function tryReduce(hand: number[]): number[] {
    if (hand.length < 4) {
        return hand
    }

    const last = hand[hand.length - 1]
    const fourth = hand[hand.length - 4]

    if (sameRank(last, fourth)) {
        return hand.slice(0, hand.length - 4)
    }
    if (sameSuit(last, fourth)) {
        const handTmp = hand.slice(0, hand.length - 3)
        return [...handTmp, hand[hand.length - 1]]
    }

    return hand
}

function sameSuit(a: card, b: card): boolean {
    return Math.floor(a/13) == Math.floor(b/13)
}

function sameRank(a: card, b: card): boolean {
    return a%13 == b%13
}

function getDeck(): deck {
    const d = []
    for (let i = 0; i < 52; i++) {
        d.push(i)
    }
    return d
}

function sum(a: number, b: number) {
    return a + b
}

function avg(ns: number[]) {
    return ns.reduce(sum, 0) / ns.length
}

function stdev(ns: number[]) {
    const s = avg(ns)
    const ss = ns.map(n => Math.pow(s - n, 2))
    return Math.sqrt(ss.reduce(sum, 0) / ns.length)
}

function main() {
    const iters = 1000000
    const results = []
    for (let i = 0; i < iters; i++) {
        results.push(playIteration())
    }

    const mean = avg(results)
    const stddev = stdev(results)
    console.log(`Mean: ${mean}\nStdev: ${stddev}`)
}

main()
