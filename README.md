<h1 align="center">neOwOfetch</h1> <p align="center">imagine neofetch, but OwO
<p align="center">
<a href="./LICENSE.md"><img src="https://cdn.discordapp.com/attachments/832652653292027904/1022794924078415932/idfk.png"></a>
<a href="https://github.com/dylanaraps/neofetch/releases"><img src="https://cdn.discordapp.com/attachments/832652653292027904/1022794924439113808/neoowowoowow.png"></a>
<a href="https://repology.org/metapackage/neofetch"><img src="https://cdn.discordapp.com/attachments/832652653292027904/1022794924783050782/bruh.png" alt="Packaging status"></a>
</p>

<img src="https://cdn.discordapp.com/attachments/917977729322872853/1022796282198237254/2022-09-23_12-36.png" alt="neofetch" align="right" height="240px">

NeOwOfetch is a shitty command-line system information tool written in `go`. NeOwOfetch displays information that you'd never want to know, software and hardware in an extremely cringe and not pleasing way.

The overall purpose of NeOwOfetch is to be used as a laughing stock. NeOwOfetch shows the information other people would have no fucking reason to see. There are so many better tools yet you decided to use this shitty bs

The information by default is displayed alongside your operating system's uwuified logo. You can further configure NeOwOfetch to instead use your shitty images or asciiarts, but ngl im too lazy and havent implemented it yet

<img src="https://cdn.discordapp.com/attachments/917977729322872853/1022797099248656445/2022-09-23_12-40.png" alt="neofetch" align="right" height="240px">

You can further configure NeOwOfetch to your horrible taste. Through the use of command-line flags and idfk i havent implemented it yet.

NeOwOfetch supports almost like 2 different operating systems. From Linux to Linux. If your favourite operating system is unsupported i dont give a shit, ill get to it when i get to it. for now its just arch and ubuntu  

<h1 align="center">how to use</h1>

after you decided that you have enough brain damage to use this mess, you can clone this repo (`git clone https://github.com/exhq/neowofetch`)  

after cloning this mess, you can either run it (`go run main.go`) or install it (`go install main.go`) which will add it to your /usr/bin


<h1 align="center">customizibility</h1>

after running the program for the first time, there should be two files in `~/.config/neowofetch/`    
### conf file

`conf` is the layout of the information. 
the syntax is `print/info */bold/bold|blue info/text`

examples:
`println italic|blue hello world` this would print an italic blue "hello world" that ends with a new line. (if you dont want the newline, replace `println` with `print`)


`info bold|yellow|nouwu GPU` this would print out your GPU in a bold yellow color without uwuifying. NOTICE: not all fonts support bold/italic

### colors file
 
this file is pretty self explainatory, you can define your own colors in RGB which you can later use in your conf file
example:
blue 0 0 255

<h1 align="center">commandline arguments</h1>

`--noascii` turns the asciiart off  
`--usepng`  uses png instead of asciiarts (still on beta)  
`--ascii=<file dir>` uses your txt file as the ascii art  
`--distro=<distroname>` forced the program to use another distro's asciiart  
`--nouwu` turns off uwuifiation for all lines  
`--nocolor` i think you can figure out what this argument does  
`--noconf` instead of using the config file, itll use a built-in default config  
`--nocolorconf` ....do i reallt have to explain the differences?  
`--16color` fallbacks to terminal's configuration instead of relying on terminal RGB support  
<p align="center">jokes aside, the asciiarts are from uwufetch, all credits go to them.</p>
