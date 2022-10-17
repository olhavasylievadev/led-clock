package main

type Placeholder [5]string

var zero = Placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = Placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = Placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = Placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = Placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = Placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = Placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = Placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = Placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = Placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var colon = Placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var dot = Placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	" ░ ",
}

//var alarm = [...]Placeholder{
//	{
//		"   ",
//		"   ",
//		"   ",
//		"   ",
//		"   ",
//	},
//	{
//		"███",
//		"█ █",
//		"███",
//		"█ █",
//		"█ █",
//	},
//	{
//		"█  ",
//		"█  ",
//		"█  ",
//		"█  ",
//		"███",
//	},
//	{
//		"███",
//		"█ █",
//		"███",
//		"█ █",
//		"█ █",
//	},
//	{
//		"██ ",
//		"█ █",
//		"██ ",
//		"█ █",
//		"█ █",
//	},
//	{
//		"█ █",
//		"███",
//		"█ █",
//		"█ █",
//		"█ █",
//	},
//	{
//		" █ ",
//		" █ ",
//		" █ ",
//		"   ",
//		" █ ",
//	},
//	{
//		"   ",
//		"   ",
//		"   ",
//		"   ",
//		"   ",
//	},
//}

var digits = [...]Placeholder{
	zero, one, two, three, four, five, six, seven, eight, nine,
}

func placeholder() {

}
