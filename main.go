package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func printWordsWithBrackets(str string) {
	fmt.Println("<----->")
	fmt.Println(str)
	fmt.Println("<----->")
}

type BadGuys interface {
	getName() string
	takeDmg(uint8)
	getHp() uint8
	getBombTime() uint8
	checkDeath() bool
	hasBomb() bool
	isFrozen() bool
	hasFakeTimer() bool
	setBomb()
	tickTock()
	shitTalk()
	freeze()
	unfreeze()
	setCloneFakeTimer()
	fakeTick()
	setHellfire()
}

type Boss struct {
	Name         string
	Hp           uint8
	BombTimer    uint8
	HellfireCD   uint8
	Bomb         bool
	Frozen       bool
	HasFakeTimer bool
	Hellfire     bool
}

func createBoss() Boss {
	return Boss{
		Name:         "Undead Wizard",
		Hp:           200,
		BombTimer:    0,
		HellfireCD:   0,
		Bomb:         false,
		Frozen:       false,
		HasFakeTimer: false,
		Hellfire:     false,
	}
}

func (b Boss) shitTalk() {
	quotes := []string{
		"You can't kill me, I'm already dead",
		"Don't expect me to throw you a bone",
		"You know dying really isn't that bad",
		"Skeletor? Never heard of him",
		"I'd rez myself, but then I'd have to pay taxes",
	}
	quote := quotes[rand.Intn(len(quotes))]
	printWordsWithBrackets(b.Name + ": " + quote)
}

func (b Boss) hasBomb() bool {
	return b.Bomb
}

func (b Boss) getBombTime() uint8 {
	return b.BombTimer
}

func (b *Boss) setBomb() {
	b.BombTimer = 4
	b.Bomb = true
}

func (b *Boss) setCloneFakeTimer() {
	b.BombTimer = 4
}

func (b *Boss) hasFakeTimer() bool {
	return b.HasFakeTimer
}

func (b *Boss) tickTock() {
	if b.BombTimer == 1 && b.Bomb {
		fmt.Println("===== BOMB EXPLODES (15 DMG) =====")
		time.Sleep(2 * time.Second)
		b.takeDmg(15)

		b.Bomb = false
		return
	}

	fmt.Println("=== TICK ===")
	time.Sleep(1200 * time.Millisecond)
	b.BombTimer--
	if b.BombTimer == 1 {
		fmt.Println(b.BombTimer, "TURN LEFT UNTIL EXPLOSION ON"+" "+b.getName())
		time.Sleep(2 * time.Second)
		return
	}
	fmt.Println(b.BombTimer, "TURNS LEFT UNTIL EXPLOSION ON"+" "+b.getName())
	time.Sleep(2 * time.Second)
}

func (b *Boss) fakeTick() {
	if b.BombTimer == 1 {
		b.HasFakeTimer = false
	} else {
		b.BombTimer--
	}
}

func (b Boss) isFrozen() bool {
	return b.Frozen
}

func (b *Boss) freeze() {
	b.Frozen = true
}

func (b *Boss) unfreeze() {
	b.Frozen = false
}

func (b Boss) getHp() uint8 {
	return b.Hp
}

func (b Boss) getName() string {
	return b.Name
}

func (b *Boss) takeDmg(dmg uint8) { // ACCOUNT FOR VOID SHIELD IN LOGIC
	if dmg >= b.Hp {
		b.Hp = 0
		b.getHp()
		return
	}
	b.Hp -= dmg
	b.getHp()
}

func (b *Boss) setHellfire() {
	b.Hellfire = true
	b.HellfireCD = 3
}

func (b Boss) checkDeath() bool {
	if b.Hp <= 0 {
		printWordsWithBrackets(b.Name + " IS DEAD AS HELL")
		time.Sleep(1650 * time.Millisecond)
		return true
	}
	return false
}

func (b *Boss) disarmBomb() {
	fmt.Println("--- BOMB DISARMED ---")
	time.Sleep(1550 * time.Millisecond)
	printWordsWithBrackets(b.Name + ": No ticking or tocking allowed")
	b.Bomb = false
	b.BombTimer = 0
}

func (b *Boss) shadowBolt(p *Player) {
	dmg := 10
	var crit bool

	critChance := rand.Intn(3)
	if critChance == 2 {
		crit = true
	}

	fmt.Println("--- " + b.Name + " CASTS SHADOW BOLT ---")
	time.Sleep(1650 * time.Millisecond)
	if crit {
		fmt.Println("--- CRIT ---")
		dmg = 15
		time.Sleep(1650 * time.Millisecond)
	}
	fmt.Println(p.getName()+" TAKES", dmg, "DMG")
}

func missileNumbers() (first, second, third uint8) {
	first, second, third = 5, 5, 5
	fcrit, scrit, tcrit := false, false, false

	for i := 0; i < 3; i++ {
		critChance := rand.Intn(9)
		if critChance == 0 {
			fcrit = true
		} else if critChance == 1 {
			scrit = true
		} else if critChance == 2 {
			tcrit = true
		}
	}

	for i := 0; i < 3; i++ {
		dmgRange := rand.Intn(18)
		if dmgRange == 0 && fcrit {
			first = 6
		} else if dmgRange == 1 && fcrit {
			first = 7
		} else if dmgRange == 2 && scrit {
			second = 6
		} else if dmgRange == 3 && scrit {
			second = 7
		} else if dmgRange == 4 && tcrit {
			third = 6
		} else if dmgRange == 5 && tcrit {
			third = 7
		}
	}

	return first, second, third
}

func (b *Boss) arcaneMissiles(p *Player) {
	first, second, third := missileNumbers()

	time.Sleep(1750 * time.Millisecond)
	fmt.Println("--- " + b.Name + " CHANNELS ARCANE MISSILES ---")
	time.Sleep(1100 * time.Millisecond)

	fmt.Println("---", first, "DMG ---")
	p.takeDmg(first)
	time.Sleep(1100 * time.Millisecond)

	fmt.Println("---", second, "DMG ---")
	p.takeDmg(second)
	time.Sleep(1100 * time.Millisecond)

	fmt.Println("---", third, "DMG ---")
	p.takeDmg(third)
	time.Sleep(1100 * time.Millisecond)
}

func (b *Boss) hellfire(p *Player) {
	fmt.Println("--- " + b.Name + " CASTS HELLFIRE ---")
	time.Sleep(1650 * time.Millisecond)
	fmt.Println(p.getName() + " IS SCORCHED FOR 7 DMG")
	p.takeDmg(7)
	p.setHellfire()
}

func (b *Boss) frozenSpike(p *Player) {
	fmt.Println("--- " + b.Name + " CASTS FROZEN SPIKE ---")
	time.Sleep(1550 * time.Millisecond)

	fmt.Println("Dodge left or right? l/r")
	var playerInput string

	for {
		fmt.Scanf("%s", playerInput)

		if playerInput == "l" {
			break
		} else if playerInput == "r" {
			break
		} else {
			continue
		}
	}

	spikeDirection := "l"
	roll := rand.Intn(2)
	if roll == 0 {
		spikeDirection = "r"
	}

	var dmg uint8 = 8
	critRoll := rand.Intn(2)

	if spikeDirection == playerInput {
		fmt.Println("- " + p.getName() + " STABBED WITH SPIKE -")
		time.Sleep(1550 * time.Millisecond)
		if critRoll == 0 {
			dmg = 12
			fmt.Println("--- CRIT ---")
			time.Sleep(1550 * time.Millisecond)
			printWordsWithBrackets(b.Name + ": OH MAN YOU GOT FUCKED UP LMAO")
		}
		p.takeDmg(dmg)
		return
	}
	fmt.Println("--- FROZEN SPIKE MISSES ---")
	time.Sleep(1550 * time.Millisecond)
	printWordsWithBrackets(b.Name + ": Ahhhh you bitch")
	time.Sleep(1350 * time.Millisecond)
	fmt.Println("--- A FLYING SHARD HITS " + b.Name + " ---")
	time.Sleep(1450 * time.Millisecond)
	printWordsWithBrackets("AH FUCK, RIGHT IN THE MANUBRIUM")
	time.Sleep(1300 * time.Millisecond)
	fmt.Println(b.Name + " TAKES 5 DMG")
	b.takeDmg(5)
}

func (b *Boss) voidShield() {
	fmt.Println("--- " + b.Name + " USES VOID SHIELD (50 ARMOR) ---")
	time.Sleep(1550 * time.Millisecond)
	printWordsWithBrackets(b.Name + ": YOU AIN'T GETTING THROUGH THAT SHIT")

	shield := uint16(b.Hp) + 50
	if shield > 250 {
		b.Hp = 250
		return
	}
	b.Hp += 50
}

var move = riggedOdds()

func riggedOdds() func() int {
	cancels := 3

	return func() int {
		for {
			roll := rand.Intn(5)
			if roll == 4 && cancels > 0 {
				cancels--
				continue
			} else if roll == 4 && cancels == 0 {
				cancels = 3
				return 4
			} else {
				return roll
			}
		}
	}
}

func (b *Boss) chooseAbility(p *Player) (spell string) {
	move := move()

	fmt.Println("<----->")

	switch move {
	case 0:
		if b.hasBomb() {
			b.disarmBomb()
			spell = ""
			return spell
		}
		b.shadowBolt(p)
		spell = "Shadow Bolt"
		return spell
	case 1:
		b.arcaneMissiles(p)
		spell = "Arcane Missiles"
		return spell
	case 2:
		b.hellfire(p)
		spell = "Hellfire"
		return spell
	case 3:
		b.frozenSpike(p)
		spell = "Frozen Spike"
		return spell
	default:
		b.voidShield()
		spell = "Void Shield"
		return spell
	}
}

type Enemy struct {
	Name         string
	Hp           uint8
	BombTimer    uint8
	PyroTimer    uint8
	Bomb         bool
	HitByShard   bool
	Frozen       bool
	Clone        bool
	PyroUsed     bool
	HasFakeTimer bool
}

func createEnemy() Enemy {
	return Enemy{
		Name:         enemyName(),
		Hp:           50,
		BombTimer:    0,
		PyroTimer:    0,
		Bomb:         false,
		HitByShard:   false,
		Frozen:       false,
		Clone:        false,
		PyroUsed:     false,
		HasFakeTimer: false,
	}
}

// CHECK IF NAME IN NAMES ARRAY OF SIZE 5
func inArr(item string, arr [5]string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return true
		}
	}
	return false
}

// CLOSURE FUNCTION FOR CREATING NAME
var enemyName = createEnemyName()

func createEnemyName() func() string { // NEXT TIME USE A SLICE AND POP NAMES ALREADY USED
	names := [5]string{
		"Goblin Bandit",
		"Troll Bandit",
		"Dwarf Bandit",
		"Elf Bandit",
		"Gnome Bandit",
	}

	var usedNames [5]string
	var index uint8

	return func() string {
		if index > 4 {
			usedNames = [5]string{}
			index = 0
		}

		var nameToReturn string

		for {
			nameToReturn = names[rand.Intn(len(names))]
			if inArr(nameToReturn, usedNames) {
				continue
			}
			usedNames[index] = nameToReturn
			index++
			break
		}

		return nameToReturn
	}
}

func (e Enemy) shitTalk() {
	hitByShardQuotes := []string{
		"That shard hurt like a mf",
		"Fucking shard came out of nowhere",
		"I still have ice in my ass from that shard",
	}
	quotes := []string{
		"Give me all your gold coins punk",
		"Hurry up, my arcane dust dealer is waiting for me",
		"You're either with us or against us",
		"Time is money, but I don't want your time, I want your money",
		"You can save 15% or more in ass whoopings by joining our gang",
	}

	quote := quotes[rand.Intn(len(quotes))]
	if e.HitByShard {
		quote = hitByShardQuotes[rand.Intn(len(quotes))]
	}

	printWordsWithBrackets(e.Name + ": " + quote)
}

func (e Enemy) hasBomb() bool {
	return e.Bomb
}

func (e Enemy) getBombTime() uint8 {
	return e.BombTimer
}

func (e *Enemy) setBomb() {
	e.BombTimer = 4
	e.Bomb = true
}

func (e *Enemy) setCloneFakeTimer() {
	e.BombTimer = 4
}

func (e *Enemy) hasFakeTimer() bool {
	return e.HasFakeTimer
}

func (e *Enemy) tickTock() {
	if e.BombTimer == 1 && e.Bomb {
		fmt.Println("===== BOMB EXPLODES (15 DMG) =====")
		time.Sleep(2 * time.Second)
		e.takeDmg(15)

		e.Bomb = false
		return
	}

	fmt.Println("=== TICK ===")
	time.Sleep(1200 * time.Millisecond)
	e.BombTimer--
	if e.BombTimer == 1 {
		fmt.Println(e.BombTimer, "TURN LEFT UNTIL EXPLOSION ON"+" "+e.getName())
		time.Sleep(2 * time.Second)
		return
	}
	fmt.Println(e.BombTimer, "TURNS LEFT UNTIL EXPLOSION ON"+" "+e.getName())
	time.Sleep(2 * time.Second)
}

func (e *Enemy) fakeTick() {
	if e.BombTimer == 1 {
		e.HasFakeTimer = false
	} else {
		e.BombTimer--
	}
}

func (e Enemy) isFrozen() bool {
	return e.Frozen
}

func (e *Enemy) freeze() {
	e.Frozen = true
}

func (e *Enemy) unfreeze() {
	e.Frozen = false
}

func (e *Enemy) setShard() {
	e.HitByShard = true
}

func (e Enemy) getHp() uint8 {
	return e.Hp
}

func (e Enemy) getName() string {
	return e.Name
}

func (e *Enemy) takeDmg(dmg uint8) {
	if dmg >= e.Hp {
		e.Hp = 0
		e.getHp()
		return
	}
	e.Hp -= dmg
	e.getHp()
}

func (e Enemy) checkDeath() bool {
	if e.Hp <= 0 {
		printWordsWithBrackets(e.Name + " IS DEAD AS HELL")
		time.Sleep(1650 * time.Millisecond)
		return true
	}
	return false
}

func (e *Enemy) useClone() {
	e.Clone = true
}

func (e *Enemy) turnCloneOff() {
	e.Clone = false
}

func (e *Enemy) cloneUsed() bool {
	return e.Clone
}

func (e *Enemy) setHellfire() {
	return // SO DARK SIM CAN STILL USE INTERFACE INSTEAD OF TYPE SWITCHES
}

// PYROBLAST
func (e *Enemy) pyroblast() {
	e.PyroUsed = true
	e.PyroTimer = 2
	fmt.Println("---", e.Name, "STARTS TO CAST PYROBLAST ---")
	time.Sleep(1750 * time.Millisecond)
}

// THROW DAGGER
func (e Enemy) throwDagger(p *Player) {
	fmt.Println("---", e.Name, "PREPARES TO THROW A DAGGER AT", p.getName(), "---")
	time.Sleep(1950 * time.Millisecond)

	var direction string = "r"
	randDirection := rand.Intn(2)
	if randDirection == 0 {
		direction = "l"
	}

	fmt.Println("Dodge left or right? l/r")
	var playerInput string

	for {
		fmt.Scanf("%s", &playerInput)
		if playerInput == "l" {
			break
		}
		if playerInput == "r" {
			break
		}
	}

	if playerInput == direction {
		fmt.Println("--- HIT ---")
		time.Sleep(1650 * time.Millisecond)
		printWordsWithBrackets(e.Name + ": BULLSEYE BITCH")
		time.Sleep(1650 * time.Millisecond)
		fmt.Println(p.getName() + ": TAKES 5 DMG")
		p.takeDmg(5)
		return
	}
	fmt.Println("--- MISS ---")
	time.Sleep(1650 * time.Millisecond)
	printWordsWithBrackets(e.Name + ": SLIPPERY BASTARD")
}

func (e *Enemy) chooseAbility(p *Player) (Spell string) {
	if e.PyroTimer > 0 {
		return "Pyroblast"
	}
	if e.PyroUsed {
		return "Pyroblast"
	}

	move := rand.Intn(3)

	var spell string

	rigged := 1

	switch move {
	case 0:
		e.pyroblast()
		spell = "Pyroblast"
		return spell
	case 1:
		e.throwDagger(p)
		spell = "Throw Dagger"
		return spell
	default:
		if rigged > 0 {
			move = rand.Intn(2)
			e.turnCloneOff()
		}

		fmt.Println("--- ENEMY CREATES CLONE ---")
		e.useClone()
		spell = "Clone"
		return spell
	}
}

type Player struct {
	Name             string
	MaxHp            uint8
	Hp               uint8
	FreezeCD         uint8
	FreezeTime       uint8
	DarkSimCD        uint8
	HellfireCD       uint8
	DoubledDown      bool
	BurnTriggered    bool
	DarkSimCloneUsed bool
	Hellfire         bool
}

func setPlayerName() string {
	var name string
	fmt.Printf("ENTER NAME: ")
	fmt.Scanln(&name)
	return name
}

func createPlayer() Player {
	return Player{
		Name:             setPlayerName(),
		MaxHp:            100,
		Hp:               100,
		FreezeCD:         0,
		FreezeTime:       0,
		DarkSimCD:        0,
		HellfireCD:       0,
		DoubledDown:      false,
		BurnTriggered:    false,
		DarkSimCloneUsed: false,
		Hellfire:         false,
	}
}

func (p Player) getName() string {
	return p.Name
}

func (p Player) getHp() uint8 {
	return p.Hp
}

func (p *Player) takeDmg(dmg uint8) {
	if dmg >= p.Hp {
		p.Hp = 0
		return
	}
	p.Hp -= dmg
}

func (p *Player) heal(amt uint8) {
	healAmt := p.Hp + amt
	if healAmt >= p.MaxHp {
		p.Hp = 100
		return
	}
	p.Hp = healAmt
}

func (p Player) checkDeath() bool {
	if p.Hp <= 0 {
		fmt.Println(p.Name, "IS DEAD AS HELL")
		return true
	}
	return false
}

func (p *Player) setHellfire() {
	p.Hellfire = true
	p.HellfireCD = 3
}

// LEECHING MINI GAMES
func crowdReactions() string {
	reactions := [8]string{
		"Crowd: 'The butt clenching, oh the butt clenching!'",
		"Crowd: 'Oh shit what's gonna happen!'",
		"Crowd: 'Oh fuck this is crazy!'",
		"Crowd: 'Mom get the camera!'",
		"Crowd: 'I just shit myself'",
		"Crowd: 'I got 50 bucks on the dumbass'",
		"Crowd: 'Holy shit this is tense'",
		"Crowd: 'Bruh ain't no way'",
	}
	return reactions[rand.Intn(8)]
}

func doorAnnouncer(correctOrNot bool) (Announcement, GoodOrBadTalk string) {
	announcements := [4]string{
		"Announcer: 'Which door will he pick?!'",
		"Announcer: 'Here he goes!'",
		"Announcer: 'The moment of truth is upon us!'",
		"Announcer: 'Will he choose the wrong door?!'",
	}
	shitTalking := [4]string{
		"Announcer: 'Damn you gotta be dumb AF to lose with those odds!'",
		"Announcer: 'And that is tragic, boy this guy sure does suck!'",
		"Announcer: 'You hate to see it, especially with those odds!'",
		"Announcer: 'My dead dad could've won this blindfolded!'",
	}
	goodTalking := [4]string{
		"Announcer: 'My God he did it!'",
		"Announcer: 'We have a winner!'",
		"Announcer: 'The skill, the elegance, the audacity!'",
		"Announcer: 'This kid is going places!'",
	}

	if correctOrNot {
		return announcements[rand.Intn(4)], goodTalking[rand.Intn(4)]
	} else {
		return announcements[rand.Intn(4)], shitTalking[rand.Intn(4)]
	}
}

// LEECHING ROCK PAPER SCISSORS
func (p *Player) leechingRPS(b BadGuys) (dmg, leech uint8, crit, win bool) {
	fmt.Println("--- LEECHING ROCK/PAPER/SCISSORS ---")
	time.Sleep(1800 * time.Millisecond)
	fmt.Println("        - BEST 2 OUT OF 3 -")
	time.Sleep(1800 * time.Millisecond)
	printWordsWithBrackets(b.getName() + ": 'I've never lost a game of rock paper scissors in my life'")
	time.Sleep(2400 * time.Millisecond)

	dmg, leech = 12, 5

	critChance := rand.Intn(4)
	if critChance != 3 {
		crit = true
		dmg += 3
		leech += 3
	}

	var cpuScore uint8 = 0
	var playerScore uint8 = 0

	for {
		fmt.Println(b.getName(), "Score:", cpuScore)
		fmt.Println(p.Name, "Score:", playerScore)
		time.Sleep(2200 * time.Millisecond)

		var cpuMove string

		cpuRoll := rand.Intn(3)
		if cpuRoll == 0 {
			cpuMove = "r"
		} else if cpuRoll == 1 {
			cpuMove = "p"
		} else {
			cpuMove = "s"
		}

		var playerMove string
		fmt.Println("CHOOSE: r/p/s")

		var cpuDrawMove string
		if cpuRoll == 0 {
			cpuDrawMove = "ROCK"
		} else if cpuRoll == 1 {
			cpuDrawMove = "PAPER"
		} else {
			cpuDrawMove = "SCISSORS"
		}

		for {
			fmt.Scanf("%v", &playerMove)
			if playerMove == "r" || playerMove == "p" || playerMove == "s" {
				time.Sleep(2200 * time.Millisecond)
				printWordsWithBrackets(crowdReactions())
				time.Sleep(2200 * time.Millisecond)
				break
			}
			continue
		}

		if cpuMove == playerMove {
			fmt.Println(b.getName(), "CHOOSES", cpuDrawMove)
			time.Sleep(1750 * time.Millisecond)
			fmt.Println("- DRAW -")
			time.Sleep(1750 * time.Millisecond)
			printWordsWithBrackets("Crowd: 'NO FUCKING WAY'")
			time.Sleep(1750 * time.Millisecond)
			continue
		} else if cpuMove == "r" && playerMove == "p" {
			fmt.Println(b.getName(), "CHOOSES ROCK")
			time.Sleep(1750 * time.Millisecond)
			fmt.Println("---", p.Name, "WINS ---")
			time.Sleep(1750 * time.Millisecond)
			printWordsWithBrackets("Crowd: 'OOOOOOOOOH FUCK'")
			time.Sleep(1750 * time.Millisecond)
			playerScore++
			if playerScore == 2 || cpuScore == 2 {
				break
			}
			continue
		} else if cpuMove == "r" && playerMove == "s" {
			fmt.Println(b.getName(), "'CHOOSES ROCK'")
			time.Sleep(1750 * time.Millisecond)
			fmt.Println("---", b.getName(), "WINS ---")
			time.Sleep(1750 * time.Millisecond)
			printWordsWithBrackets("Crowd: 'OOOOOOOOOH SHIT'")
			time.Sleep(1750 * time.Millisecond)
			cpuScore++
			if playerScore == 2 || cpuScore == 2 {
				break
			}
			continue
		} else if cpuMove == "p" && playerMove == "r" {
			fmt.Println(b.getName(), "CHOOSES PAPER")
			time.Sleep(1750 * time.Millisecond)
			fmt.Println("---", b.getName(), "WINS ---")
			time.Sleep(1750 * time.Millisecond)
			printWordsWithBrackets("Crowd: 'OOOOOOOOOH DAMN'")
			time.Sleep(1750 * time.Millisecond)
			cpuScore++
			if playerScore == 2 || cpuScore == 2 {
				break
			}
			continue
		} else if cpuMove == "p" && playerMove == "s" {
			fmt.Println(b.getName(), "CHOOSES PAPER")
			time.Sleep(1750 * time.Millisecond)
			fmt.Println("---", p.Name, "WINS ---")
			time.Sleep(1750 * time.Millisecond)
			printWordsWithBrackets("Crowd: 'OOOOOOOOOH FUCK'")
			time.Sleep(1750 * time.Millisecond)
			playerScore++
			if playerScore == 2 || cpuScore == 2 {
				break
			}
			continue
		} else if cpuMove == "s" && playerMove == "r" {
			fmt.Println(b.getName(), "CHOOSES SCISSORS")
			time.Sleep(1750 * time.Millisecond)
			fmt.Println("---", p.Name, "WINS ---")
			time.Sleep(1750 * time.Millisecond)
			printWordsWithBrackets("Crowd: 'OOOOOOOOOH SHIT'")
			time.Sleep(1750 * time.Millisecond)
			playerScore++
			if playerScore == 2 || cpuScore == 2 {
				break
			}
			continue
		} else if cpuMove == "s" && playerMove == "p" {
			fmt.Println(b.getName(), "CHOOSES SCISSORS")
			time.Sleep(1750 * time.Millisecond)
			fmt.Println("---", b.getName(), "WINS ---")
			time.Sleep(1750 * time.Millisecond)
			printWordsWithBrackets("Crowd: 'OOOOOOOOOH DAMN'")
			time.Sleep(1750 * time.Millisecond)
			cpuScore++
			if playerScore == 2 || cpuScore == 2 {
				break
			}
		}
	}

	if playerScore == 2 {
		win = true
		printWordsWithBrackets(b.getName() + ": You got lucky punk")
		time.Sleep(1650 * time.Millisecond)
	} else {
		printWordsWithBrackets(b.getName() + ": All skill, get shrekt")
		time.Sleep(1650 * time.Millisecond)
	}

	return dmg, leech, crit, win
}

// LEECHING COIN FLIP
func (p *Player) leechingCoinFlip(b BadGuys) (dmg, leech uint8, crit, win bool) {
	dmg, leech = 9, 4

	coinFlips := [5]string{}

	for i := 0; i < 5; i++ {
		coinFlip := rand.Intn(2)
		if coinFlip == 0 {
			coinFlips[i] = "HEADS"
		} else {
			coinFlips[i] = "TAILS"
		}
	}

	var playerChoice string
	fmt.Println("--- HEADS OR TAILS ---")
	time.Sleep(2200 * time.Millisecond)
	fmt.Println("h/t ?")

	for {
		fmt.Scanf("%v", &playerChoice)
		if playerChoice != "h" && playerChoice != "t" {
			continue
		}
		break
	}

	var playerFlip string
	if playerChoice == "h" {
		playerFlip = "HEADS"
	} else {
		playerFlip = "TAILS"
	}

	for i := 0; i < 5; i++ {
		if i == 4 {
			fmt.Println("FINAL FLIP")
			time.Sleep(1650 * time.Millisecond)
			fmt.Println(coinFlips[i])
			time.Sleep(1850 * time.Millisecond)
			break
		}
		fmt.Println(coinFlips[i])
		time.Sleep(1650 * time.Millisecond)
		printWordsWithBrackets(crowdReactions())
		time.Sleep(1650 * time.Millisecond)
	}

	critChance := rand.Intn(2)

	if playerFlip == coinFlips[4] {
		printWordsWithBrackets("YOU WIN")
		time.Sleep(1 * time.Second)
		printWordsWithBrackets(b.getName() + ": Son of a bitch")
		time.Sleep(1400 * time.Millisecond)
		if critChance == 0 {
			crit = true
			win = true
			dmg += 3
			leech += 3
		} else {
			win = true
		}
	} else {
		printWordsWithBrackets("YOU LOSE")
		time.Sleep(1 * time.Second)
		printWordsWithBrackets(b.getName() + ": Ya boiiiiiiiiii")
	}

	return dmg, leech, crit, win
}

// LEECHING DOORS
func (p *Player) leechingDoors() (dmg, leech uint8, crit, win bool) {
	dmg, leech = 7, 3

	wrongDoor := rand.Intn(4) + 1

	fmt.Println("--- LEECHING DOORS ---")
	printWordsWithBrackets("Announcer: There are 4 doors, one of which has a bomb behind it, enjoy the show folks!")
	time.Sleep(2400 * time.Millisecond)
	fmt.Println("    _       _       _       _")
	fmt.Println("   | |     | |     | |     | |")
	fmt.Println("1. |_|  2. |_|  3. |_|  4. |_|")
	time.Sleep(1200 * time.Millisecond)

	preChoice, _ := doorAnnouncer(false)
	printWordsWithBrackets(preChoice)

	var playerChoice int
	for {
		fmt.Scanf("%d", &playerChoice)
		if playerChoice != 1 && playerChoice != 2 && playerChoice != 3 && playerChoice != 4 {
			continue
		}
		break
	}

	if playerChoice != wrongDoor {
		fmt.Println("--- WINNER ---")
		win = true
		time.Sleep(1500 * time.Millisecond)
	} else {
		fmt.Println("--- BOMB EXPLODES ---")
		time.Sleep(1650 * time.Millisecond)
	}

	critRoll := rand.Intn(3)
	if critRoll == 0 {
		crit = true
	}

	_, quote := doorAnnouncer(win)
	printWordsWithBrackets(quote)
	time.Sleep(1950 * time.Millisecond)

	return dmg, leech, crit, win
}

// BASE FUNCTION
func (p *Player) leechingGames(b BadGuys) {
	game := rand.Intn(3)
	var dmg, leech uint8
	var crit, win bool

	if game == 0 {
		dmg, leech, crit, win = p.leechingRPS(b)
	} else if game == 1 {
		dmg, leech, crit, win = p.leechingCoinFlip(b)
	} else {
		dmg, leech, crit, win = p.leechingDoors()
	}

	doubleDown := false
	if win && crit {
		doubleDown = true
		fmt.Println("--- CRIT ---")
		time.Sleep(1700 * time.Millisecond)
		fmt.Println(b.getName() + ": Oh this is bullshit")
		time.Sleep(1500 * time.Millisecond)
	}

	var playerChoice string

	if doubleDown {
		printWordsWithBrackets("                   *** DOUBLE DOWN? ***\nDOUBLE DMG AND LEECH IF YOU WIN, TAKE ORIGINAL DMG IF YOU LOSE y/n")
		fmt.Println("         (Original DMG:", dmg, " Original Leech:", strconv.Itoa(int(leech))+")")
		doubleWin := false

		for {
			fmt.Scanf("%v", &playerChoice)
			if playerChoice == "y" && game == 0 {
				_, _, _, doubleWin = p.leechingRPS(b)
				if doubleWin {
					b.takeDmg(dmg * 2)
					fmt.Println(b.getName(), "TAKES", dmg*2, "DMG, FUCK")
					p.heal(leech * 2)
					fmt.Println(p.Name, "HEALS FOR", leech*2, "HP")
				} else {
					p.takeDmg(dmg)
					fmt.Println(p.Name, "TAKES", dmg, "DMG, DAMN")
				}
				break
			} else if playerChoice == "y" && game == 1 {
				_, _, _, doubleWin = p.leechingCoinFlip(b)
				if doubleWin {
					b.takeDmg(dmg * 2)
					fmt.Println(b.getName(), "TAKES", dmg*2, "DMG, FUCK")
					p.heal(leech * 2)
					fmt.Println(p.Name, "HEALS FOR", leech*2, "HP")
				} else {
					p.takeDmg(dmg)
					fmt.Println(p.Name, "TAKES", dmg, "DMG, DAMN")
				}
				break
			} else if playerChoice == "y" && game == 2 {
				_, _, _, doubleWin = p.leechingDoors()
				if doubleWin {
					b.takeDmg(dmg * 2)
					fmt.Println(b.getName(), "TAKES", dmg*2, "DMG, FUCK")
					p.heal(leech * 2)
					fmt.Println(p.Name, "HEALS FOR", leech*2, "HP")
				} else {
					p.takeDmg(dmg)
					fmt.Println(p.Name, "TAKES", dmg, "DMG, DAMN")
				}
				break
			} else if playerChoice == "n" {
				b.takeDmg(dmg)
				fmt.Println(b.getName(), "TAKES", dmg, "DMG")
				p.heal(leech)
				fmt.Println(p.Name, "HEALS FOR", leech, "HP")
				break
			}
		}
	} else if win && !doubleDown {
		b.takeDmg(dmg)
		fmt.Println(b.getName(), "TAKES", dmg, "DMG")
		p.heal(leech)
		fmt.Println(p.Name, "HEALS FOR", leech, "HP")
	} else if game == 0 && !win {
		p.takeDmg(3)
		fmt.Println(p.Name, "TAKES 3 DMG")
	} else if game == 1 && !win {
		p.takeDmg(2)
		fmt.Println(p.Name, "TAKES 2 DMG")
	} else if game == 2 && !win {
		p.takeDmg(5)
		fmt.Println(p.Name, "TAKES 5 DMG")
	}
}

// FLAMING ORB
func (p *Player) flamingOrb(l uint8, b BadGuys) {
	var dmg uint8 = 3

	if l == 1 {
		fmt.Println("--- FLAMING ORB HITS", b.getName(), "FOR", dmg, "DMG ---")
		b.takeDmg(dmg)
		fmt.Println(b.getName() + ": HAH, that's it?!")
		time.Sleep(1150 * time.Millisecond)
		if b.hasBomb() {
			b.tickTock()
			fmt.Println(b.getName() + ": Oh fuck")
		}
		time.Sleep(1650 * time.Millisecond)
		return
	}
	dmg = l * 3
	fmt.Println("--- FLAMING ORB HITS", b.getName(), "FOR", dmg, "DMG ---")
	b.takeDmg(dmg)
	time.Sleep(1450 * time.Millisecond)
	fmt.Println("***THE FLAME SPREADS***")
	time.Sleep(1350 * time.Millisecond)
	p.BurnTriggered = true
}

// FREEZE
func (p *Player) freeze(b BadGuys) {
	if _, ok := b.(*Boss); ok {
		fmt.Println(b.getName() + " IS FROZEN FOR 2 TURNS")
		time.Sleep(1750 * time.Millisecond)
		b.freeze()
		p.FreezeTime = 2
		p.FreezeCD = 5
		return
	}
	fmt.Println(b.getName() + " IS FROZEN FOR 1 TURN")
	time.Sleep(1750 * time.Millisecond)
	b.freeze()
	p.FreezeTime = 1
	p.FreezeCD = 4
}

// DARK SIMULACRUM

// PYROBLAST COPY
func (p *Player) pyroblastDarkSim(b BadGuys) {
	printWordsWithBrackets(p.Name + ": Nice fireball scrub, check this shit out")
	time.Sleep(1850 * time.Millisecond)
	fmt.Println("---", p.Name, "STARTS TO CAST PYROBLAST ---")
	time.Sleep(1950 * time.Millisecond)
	printWordsWithBrackets(b.getName() + ": Spell stealing bastard")
	time.Sleep(1850 * time.Millisecond)
	fmt.Println("--- " + b.getName() + ": HIT FOR 20 DMG ---")
	b.takeDmg(20)
	time.Sleep(1750 * time.Millisecond)
}

// THROW DAGGER COPY
func (p *Player) throwDaggerDarkSim(b BadGuys) {
	printWordsWithBrackets(p.Name + ": THAT'S NOT A KNIFE, THIS IS A KNIFE")
	time.Sleep(1650 * time.Millisecond)
	fmt.Println("---", p.Name, "PREPARES TO THROW A DAGGER AT", b.getName(), "---")
	time.Sleep(1950 * time.Millisecond)

	var direction string = "RIGHT"
	randDirection := rand.Intn(2)
	if randDirection == 0 {
		direction = "LEFT"
	}

	fmt.Println("Dodge left or right? l/r")
	time.Sleep(1550 * time.Millisecond)
	printWordsWithBrackets(fmt.Sprintf("%s: AH SHIT, %s!", b.getName(), direction))
	time.Sleep(1550 * time.Millisecond)

	if direction == "RIGHT" {
		fmt.Println("--- HIT ---")
		time.Sleep(1650 * time.Millisecond)
		printWordsWithBrackets(p.Name + ": BULLSEYE BITCH MUAHAHA")
		time.Sleep(1650 * time.Millisecond)
		fmt.Println(b.getName() + ": TAKES 15 DMG")
		b.takeDmg(15)
		return
	}
	fmt.Println("--- MISS ---")
	time.Sleep(1650 * time.Millisecond)
	printWordsWithBrackets(p.Name + ": SLIPPERY BASTARD")
}

// CLONE COPY
func (p *Player) cloneDarkSim(b BadGuys) {
	p.DarkSimCloneUsed = true

	printWordsWithBrackets(p.Name + ": Oh you're not gonna like this")
	time.Sleep(1650 * time.Millisecond)
	fmt.Println("---", p.Name, "SPLITS IN TWO ---")
	time.Sleep(1650 * time.Millisecond)
	fmt.Println(p.Name + ": WHICH ONE IS THE REAL ME?")
	fmt.Println(p.Name + ": WHICH ONE IS THE REAL ME?")
	time.Sleep(1850 * time.Millisecond)
	printWordsWithBrackets(b.getName() + ": OH SHIT")

	guess := rand.Intn(2)
	clone := rand.Intn(2)

	if guess == 0 {
		printWordsWithBrackets(b.getName() + ": THROWS KNIFE AT LEFT CLONE")
		time.Sleep(1650 * time.Millisecond)
	} else {
		printWordsWithBrackets(b.getName() + ": THROWS KNIFE AT RIGHT CLONE")
		time.Sleep(1650 * time.Millisecond)
	}

	if guess == clone {
		fmt.Println("--- CLONE SHATTERS ---")
		time.Sleep(1500 * time.Millisecond)
		printWordsWithBrackets(b.getName() + ": OH FUCK")
		time.Sleep(1550 * time.Millisecond)
		fmt.Println("---", p.Name, "APPEARS BEHIND", b.getName(), "---")
		time.Sleep(1600 * time.Millisecond)
		fmt.Println("--- SHANK 30 DMG ---")
		b.takeDmg(30)
		time.Sleep(1650 * time.Millisecond)
	} else {
		fmt.Println("---", p.Name, "HIT FOR 5 DMG ---")
		time.Sleep(1500 * time.Millisecond)
		printWordsWithBrackets(b.getName() + ": TAKE THAT YOU BITCH")
		p.takeDmg(5)
		time.Sleep(1650 * time.Millisecond)
	}
}

// SHADOW BOLT COPY
func (p *Player) shadowBoltDarkSim(b BadGuys) {
	printWordsWithBrackets(p.Name + ": I'LL BE TAKING THAT SIR")
	time.Sleep(1650 * time.Millisecond)
	printWordsWithBrackets(p.Name + " STARTS CASTING SHADOW BOLT")

	dmg := 15

	crit := false
	critChance := rand.Intn(3)
	if critChance == 0 {
		crit = true
		dmg = 25
	}

	if crit {
		fmt.Println("--- CRIT ---")
		time.Sleep(1450 * time.Millisecond)
		printWordsWithBrackets(b.getName() + ": AND IT CRIT?!")
		time.Sleep(1350 * time.Millisecond)
	}
	fmt.Println("--- SHADOW BOLT HITS", b.getName()+" FOR", dmg, "DMG", "---")
}

// ARCANE MISSILES HELPER FUNCTION
func missileNumbersCopy() (first, second, third uint8) {
	first, second, third = 6, 6, 6
	fcrit, scrit, tcrit := false, false, false

	for i := 0; i < 3; i++ {
		critChance := rand.Intn(9)
		if critChance == 0 {
			fcrit = true
		} else if critChance == 1 {
			scrit = true
		} else if critChance == 2 {
			tcrit = true
		}
	}

	for i := 0; i < 3; i++ {
		dmgRange := rand.Intn(18)
		if dmgRange == 0 && fcrit {
			first = 8
		} else if dmgRange == 1 && fcrit {
			first = 9
		} else if dmgRange == 2 && scrit {
			second = 8
		} else if dmgRange == 3 && scrit {
			second = 9
		} else if dmgRange == 4 && tcrit {
			third = 8
		} else if dmgRange == 5 && tcrit {
			third = 9
		}
	}

	return first, second, third
}

// ARCANE MISSILES COPY
func (p *Player) arcaneMissilesDarkSim(b BadGuys) {
	first, second, third := missileNumbersCopy()

	printWordsWithBrackets(p.Name + ": This is how the pros do it")
	time.Sleep(1750 * time.Millisecond)
	fmt.Println("--- " + p.Name + " CHANNELS ARCANE MISSILES ---")
	time.Sleep(1450 * time.Millisecond)

	fmt.Println("---", first, "DMG ---")
	b.takeDmg(first)
	time.Sleep(1650 * time.Millisecond)

	fmt.Println("---", second, "DMG ---")
	b.takeDmg(second)
	time.Sleep(1650 * time.Millisecond)

	fmt.Println("---", third, "DMG ---")
	b.takeDmg(third)
	time.Sleep(1650 * time.Millisecond)
}

// HELLFIRE COPY
func (p *Player) hellfireDarkSim(b BadGuys) {
	printWordsWithBrackets(p.Name + ": My turn to roast your ass")
	time.Sleep(1650 * time.Millisecond)
	fmt.Println("--- " + p.Name + " CASTS HELLFIRE ---")
	time.Sleep(1650 * time.Millisecond)
	fmt.Println(b.getName() + " IS SCORCHED FOR 10 DMG")
	b.takeDmg(10)
	b.setHellfire()
}

// FROZEN SPIKE COPY
func (p *Player) frozenSpikeDarkSim(b BadGuys) {
	printWordsWithBrackets(p.Name + " Think fast!")
	time.Sleep(1450 * time.Millisecond)
	fmt.Println("Dodge left or right? l/r")
	time.Sleep(1450 * time.Millisecond)

	printWordsWithBrackets(b.getName() + ": AH, FUCK . . .")

	spikeDirection := "LEFT"
	roll := rand.Intn(2)
	if roll == 0 {
		spikeDirection = "RIGHT"
	}

	var dmg uint8 = 12
	critRoll := rand.Intn(2)

	if spikeDirection == "LEFT" {
		fmt.Println("- " + b.getName() + " STABBED WITH SPIKE -")
		time.Sleep(1550 * time.Millisecond)
		if critRoll == 0 {
			dmg = 18
			fmt.Println("--- CRIT ---")
			time.Sleep(1550 * time.Millisecond)
			printWordsWithBrackets(b.getName() + ": You gotta be kidding me")
		}
		p.takeDmg(dmg)
		return
	}
	fmt.Println("--- FROZEN SPIKE MISSES ---")
	time.Sleep(1550 * time.Millisecond)
	printWordsWithBrackets(b.getName() + ": HA, YOU MISSED DUMBA-")
	time.Sleep(1100 * time.Millisecond)
	fmt.Println("--- A FLYING SHARD HITS " + b.getName() + " IN THE ASS ---")
	time.Sleep(1650 * time.Millisecond)
	printWordsWithBrackets("GOD DAMN, FUCK, SHIT")
	time.Sleep(1300 * time.Millisecond)
	fmt.Println(b.getName() + " TAKES 5 DMG")
	b.takeDmg(5)
}

// VOID SHIELD COPY
func (p *Player) voidShieldDarkSim(b BadGuys) {
	printWordsWithBrackets(p.Name + ": Nice spell, let me try it out")
	time.Sleep(1650 * time.Millisecond)
	fmt.Println("--- VOID SHIELD ACTIVATED ---")
	time.Sleep(1650 * time.Millisecond)

	printWordsWithBrackets(p.Name + ": Let me see if I can improve this thing")
	time.Sleep(1650 * time.Millisecond)

	fmt.Println(p.Name + " ATTEMPTS TO ENCHANT VOID SHIELD")
	time.Sleep(1650 * time.Millisecond)

	flip := rand.Intn(2)
	if flip == 0 {
		fmt.Println("- SUCCES -")
		time.Sleep(1 * time.Second)
		fmt.Println("- VOID SHIELD ENCHANTED - (+25 ARMOR)")
		time.Sleep(1450 * time.Millisecond)
		printWordsWithBrackets(b.getName() + ": This is some bullshit")

		shield := uint16(p.Hp) + 75
		if shield > 250 {
			p.Hp = 250
			return
		}
		p.Hp += 75
		return
	}
	fmt.Println("- FAIL -")
	time.Sleep(1 * time.Second)
	printWordsWithBrackets(b.getName() + ": Just a normal shield huh?")

	shield := uint16(p.Hp) + 50
	if shield > 250 {
		p.Hp = 250
		return
	}
	p.Hp += 50
}

// BASE FUNCTION
func (p *Player) darkSimulacrum(b BadGuys, currentSpell string) {
	p.DarkSimCD = 4

	if currentSpell == "Pyroblast" {
		p.pyroblastDarkSim(b)
	} else if currentSpell == "Throw Dagger" {
		p.throwDaggerDarkSim(b)
	} else if currentSpell == "Clone" {
		p.cloneDarkSim(b)
	} else if currentSpell == "Shadow Bolt" {
		p.shadowBoltDarkSim(b)
	} else if currentSpell == "Arcane Missiles" {
		p.arcaneMissilesDarkSim(b)
	} else if currentSpell == "Hellfire" {
		p.hellfireDarkSim(b)
	} else if currentSpell == "Frozen Spike" {
		p.frozenSpikeDarkSim(b)
	} else if currentSpell == "Void Shield" {
		p.voidShieldDarkSim(b)
	}
}

func (p *Player) chooseAbility(c uint8, b BadGuys, currentSpell string, wasEnemyClone bool) {
	if b.isFrozen() {
		fmt.Println(b.getName() + ": OH SHIT (FROZEN)")
		time.Sleep(2400 * time.Millisecond)
	} else {
		b.shitTalk()
		time.Sleep(2400 * time.Millisecond)
	}

	var playerInput uint8

	timeBomb := "1: TimeBomb"
	if b.hasBomb() {
		timeBomb = "1: TimeBomb(ON COOLDOWN " + strconv.Itoa(int(b.getBombTime())) + ")"
	}

	freezeSpell := "4: Freeze"
	if p.FreezeCD > 0 {
		freezeSpell = "4: Freeze(ON COOLDOWN " + strconv.Itoa(int(p.FreezeCD)) + ")"
	}

	enemyAbilityCopy := currentSpell
	var darkSim string

	// SHOW DARK SIM SPELL IF NOT ON CD
	if enemyAbilityCopy != "" && p.DarkSimCD == 0 {
		darkSim = "5: Dark Simulacrum(" + enemyAbilityCopy + ")"
	}
	// SHOW DARK SIM NO ENEMY SPELL USED
	if enemyAbilityCopy == "" {
		darkSim = "5: Dark Simulacrum(NO ENEMY ABILITY TO STEAL)"
	}
	// SHOW DARK SIM ON CD
	if p.DarkSimCD > 0 {
		darkSim = "5: Dark Simulacrum(ON COOLDOWN " + strconv.Itoa(int(p.DarkSimCD)) + ")"
	}

	fmt.Println("----------------------")
	fmt.Println(p.getName()+" HP:", p.getHp())
	fmt.Println(b.getName()+" HP:", b.getHp())
	fmt.Println("----------------------")
	fmt.Println(timeBomb)
	fmt.Println("2: Leeching Games")
	fmt.Println("3: Flaming Orb")
	fmt.Println(freezeSpell)
	fmt.Println(darkSim)

	for {
		fmt.Scan(&playerInput)

		if playerInput == 1 {
			if b.hasBomb() || b.hasFakeTimer() {
				fmt.Println("BOMB ALREADY ACTIVE")
				continue
			}
			// CLONE LOGIC
			if wasEnemyClone {
				fmt.Println("- CLONE HIT -")
				time.Sleep(1450 * time.Millisecond)
				fmt.Println("--- CLONE EXPLODES ---")
				b.setCloneFakeTimer()
				time.Sleep(1450 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": GLAD THAT WASN'T ME LMAO")
				time.Sleep(1650 * time.Millisecond)
				break
			}
			if !wasEnemyClone && currentSpell == "Clone" {
				fmt.Println("---", b.getName()+" HIT", "---")
				time.Sleep(1650 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": AH FUCK")
				time.Sleep(1550 * time.Millisecond)
				fmt.Println("--- CLONE VANISHES ---")
				time.Sleep(1300 * time.Millisecond)
			}
			fmt.Println("===== BOMB SET ON", b.getName(), "=====")
			time.Sleep(2 * time.Second)
			b.setBomb()
			break
		} else if playerInput == 2 {
			// CLONE LOGIC
			if wasEnemyClone {
				fmt.Println("- CLONE HIT -")
				time.Sleep(1450 * time.Millisecond)
				fmt.Println("--- CLONE EXPLODES ---")
				time.Sleep(1650 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": I AIN'T PLAYING NO GAMES BOI")
				time.Sleep(1650 * time.Millisecond)
				break
			}
			if !wasEnemyClone && currentSpell == "Clone" {
				fmt.Println("---", b.getName()+" HIT", "---")
				time.Sleep(1650 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": AH FUCK")
				time.Sleep(1550 * time.Millisecond)
				fmt.Println("--- CLONE VANISHES ---")
				time.Sleep(1300 * time.Millisecond)
			}
			p.leechingGames(b)
			break
		} else if playerInput == 3 {
			// CLONE LOGIC
			if wasEnemyClone {
				fmt.Println("- CLONE HIT -")
				time.Sleep(1450 * time.Millisecond)
				fmt.Println("--- FLAMING ORB TORCHES CLONE ---")
				time.Sleep(1650 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": DAMN, THAT COULD'VE BEEN ME")
				time.Sleep(1650 * time.Millisecond)
				break
			}
			if !wasEnemyClone && currentSpell == "Clone" {
				fmt.Println("---", b.getName()+" HIT", "---")
				time.Sleep(1650 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": AH FUCK")
				time.Sleep(1550 * time.Millisecond)
				fmt.Println("--- CLONE VANISHES ---")
				time.Sleep(1300 * time.Millisecond)
			}
			p.flamingOrb(c, b)
			break
		} else if playerInput == 4 {
			if p.FreezeCD > 0 {
				continue
			}
			// CLONE LOGIC
			if wasEnemyClone {
				fmt.Println("- CLONE HIT -")
				time.Sleep(1450 * time.Millisecond)
				fmt.Println("--- CLONE SHATTERS ---")
				time.Sleep(1550 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": THAT'S JUST COLD")
				time.Sleep(1650 * time.Millisecond)
				p.FreezeCD = 4
				break
			}
			if !wasEnemyClone && currentSpell == "Clone" {
				fmt.Println("---", b.getName()+" HIT", "---")
				time.Sleep(1650 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": AH FUCK")
				time.Sleep(1550 * time.Millisecond)
				fmt.Println("--- CLONE VANISHES ---")
				time.Sleep(1300 * time.Millisecond)
			}
			p.freeze(b)
			break
		} else if playerInput == 5 {
			if p.DarkSimCD > 0 || enemyAbilityCopy == "" {
				continue
			}
			// CLONE LOGIC
			if wasEnemyClone {
				fmt.Println("- CLONE HIT -")
				time.Sleep(1450 * time.Millisecond)
				fmt.Println("--- CLONE MELTS ---")
				time.Sleep(1650 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": NO COPYING MF")
				time.Sleep(1650 * time.Millisecond)
				break
			}
			if !wasEnemyClone && currentSpell == "Clone" {
				fmt.Println("---", b.getName()+" HIT", "---")
				time.Sleep(1650 * time.Millisecond)
				printWordsWithBrackets(b.getName() + ": AH FUCK")
				time.Sleep(1550 * time.Millisecond)
				fmt.Println("--- CLONE VANISHES ---")
				time.Sleep(1300 * time.Millisecond)
			}
			p.darkSimulacrum(b, currentSpell)
			break
		} else {
			fmt.Println("FUCKBOI, PICK AN ABILITY")
			continue
		}
	}
}

// HANDLERS
func handleTimeBomb(b BadGuys) {
	if b.hasBomb() && !b.hasFakeTimer() {
		b.tickTock()
	}
	if b.hasFakeTimer() {
		b.fakeTick()
	}
}

func (p *Player) handleBurn(f func()) {
	if p.BurnTriggered {
		f()
		p.BurnTriggered = false
	}
}

func burn(e []Enemy) func() {
	flag := false

	return func() {
		flag = true
		if flag {
			for i := 1; i < len(e); i++ {
				e[i].takeDmg(3)
				fmt.Println(e[i].getName(), "TAKES 3 DMG")
				time.Sleep(1200 * time.Millisecond)
			}
			flag = false
		}
	}
}

func (p *Player) handleFreeze(b BadGuys) {
	if p.FreezeTime == 0 {
		b.unfreeze()
		if p.FreezeCD != 0 {
			p.FreezeCD--
		}
		return
	}
	p.FreezeTime--
	p.FreezeCD--
}

func (p *Player) handleDarkSimulacrum() {
	if p.DarkSimCD > 0 {
		p.DarkSimCD--
	}
}

func (e *Enemy) handlePyroblast(p *Player) {
	if e.PyroTimer == 0 && !e.Frozen && e.PyroUsed {
		fmt.Println("--- PYROBLAST HITS", p.getName(), "FOR 10 DMG ---")
		p.takeDmg(10)
		e.PyroUsed = false
		return
	}
	if e.PyroTimer > 0 && !e.Frozen {
		fmt.Println("-", e.PyroTimer, "TURNS UNTIL", p.getName(), "IS HIT BY PYROBLAST -")
		e.PyroTimer--
	}
}

func (e *Enemy) handleClone() (WasClone bool) {
	e.turnCloneOff()

	fmt.Println("WHICH ENEMY IS THE REAL ONE? l/r")
	var guess string

	for {
		fmt.Scanf("%s", &guess)
		if guess == "l" {
			break
		} else if guess == "r" {
			break
		} else {
			continue
		}
	}

	realClone := rand.Intn(2)
	var cloneString string
	if realClone == 0 {
		cloneString = "l"
	} else {
		cloneString = "r"
	}

	return guess == cloneString
}

func (p *Player) handleHellfireDot() {
	if p.Hellfire {
		if p.HellfireCD == 0 {
			p.Hellfire = false
			return
		}
		fmt.Println("- " + p.Name + ": CONTINUES TO BURN FOR 3 DMG -")
		p.takeDmg(3)
		p.HellfireCD--
	}
}

func (b *Boss) handleHellfireDotDarkSim() {
	if b.Hellfire {
		if b.HellfireCD == 0 {
			b.Hellfire = false
			return
		}
		fmt.Println("- " + b.Name + ": CONTINUES TO BURN FOR 5 DMG -")
		b.takeDmg(5)
		b.HellfireCD--
	}
}

// TRACK ENEMIES LEFT IN OUTER LOOP
func enemyCount(e []Enemy) uint8 {
	var count uint8
	count = uint8(len(e))
	return count
}

func main() {
	/*var enemies [5]Enemy

	for i := 0; i < 5; i++ {
		enemies[i] = createEnemy()
	}

	for i := 0; i < len(enemies); i++ {
		fmt.Println(enemies[i].getName() + " SPAWNS")
		time.Sleep(1150 * time.Millisecond)
	}*/

	p := createPlayer()

	// NORMAL ENEMIES GAME LOOP
	/*for i := 0; i < len(enemies); i++ {

		// CURRENT ENEMY COUNT
		ec := enemyCount(enemies[i:])

		// CLOSURE FOR AOE BURN DMG
		aoeBurn := burn(enemies[i:])

		// DARKSIM COPY
		var currentSpell string

		for {
			var wasClone bool

			if enemies[i].cloneUsed() {
				wasClone = enemies[i].handleClone()
			}

			// PLAYER TURN
			p.chooseAbility(ec, &enemies[i], currentSpell, wasClone)
			p.handleBurn(aoeBurn)
			handleTimeBomb(&enemies[i])
			p.handleFreeze(&enemies[i])
			p.handleDarkSimulacrum()

			// DEATH CHECKS
			if enemies[i].checkDeath() {
				break
			}
			if p.checkDeath() {
				log.Fatal("===== YOU LOSE =====")
			}

			// SKIP ENEMY TURN IF FROZEN
			if enemies[i].isFrozen() {
				currentSpell = ""
				continue
			}

			// SKIP ENEMY TURN IF DARK SIM SUCCESSFULLY USED ON CLONE
			if p.DarkSimCloneUsed {
				p.DarkSimCloneUsed = false
				currentSpell = ""
				wasClone = false
				continue
			}

			// ENEMY TURN
			currentSpell = enemies[i].chooseAbility(&p)
			enemies[i].handlePyroblast(&p)

			// DEATH CHECKS
			if p.checkDeath() {
				log.Fatal("===== YOU LOSE =====")
			}
			if enemies[i].checkDeath() {
				break
			}

			time.Sleep(1950 * time.Millisecond)
		}
	}*/

	fmt.Println("=== ALL ENEMIES DEFEATED ===")
	time.Sleep(1550 * time.Millisecond)

	fmt.Println("*** HEALED FOR FULL HP ***")
	p.heal(100)
	time.Sleep(2200 * time.Millisecond)

	b := createBoss()

	// BOSS GAME LOOP
	for {
		// BOSS COUNT 1 FOR FLAMING ORB
		var ec uint8 = 1

		// DARKSIM COPY
		var currentSpell string

		for {
			// PLAYER TURN
			p.handleHellfireDot()
			p.chooseAbility(ec, &b, currentSpell, false)
			handleTimeBomb(&b)
			p.handleFreeze(&b)
			p.handleDarkSimulacrum()

			// DEATH CHECKS
			if b.checkDeath() {
				break
			}
			if p.checkDeath() {
				log.Fatal("===== YOU LOSE =====")
			}

			// SKIP BOSS TURN IF FROZEN
			if b.isFrozen() {
				currentSpell = ""
				continue
			}

			// BOSS TURN
			b.handleHellfireDotDarkSim()
			currentSpell = b.chooseAbility(&p)

			// DEATH CHECKS
			if p.checkDeath() {
				log.Fatal("===== YOU LOSE =====")
			}
			if b.checkDeath() {
				break
			}

			time.Sleep(1950 * time.Millisecond)
		}
		break
	}

	fmt.Println("===== YOU WIN =====")
}
