package main

import (
	"fmt"
	"os"
	"os/exec"
	"io"
	"bufio"
)

import flag "github.com/spf13/pflag"

type selpg_args struct {
	start_page int
	end_page int
	in_filename string
	page_len int
	page_type bool
	print_dest string
}

func check(sa selpg_args){
	if (sa.start_page < 0){
		fmt.Fprintf(os.Stderr, "start_page should be specified correctly!\n")
		os.Exit(1)
	}
	if (sa.end_page < 0){
		fmt.Fprintf(os.Stderr, "end_page should be specified correctly!\n")
		os.Exit(1)
	}
	if(sa.start_page>sa.end_page){
		fmt.Fprintf(os.Stderr, "start_page should be less than end_page!\n")
		os.Exit(1)
	}
	if(sa.page_type == true) && (sa.page_len!=72){
		fmt.Fprintf(os.Stderr, "-l and -lf is exclusive\n")
		os.Exit(1)
	}
}

func main(){
	var sa selpg_args;
    flag.IntVarP(&(sa.start_page), "start_page", "s",-1, "specify page of start.  defaults to -1.")
    flag.IntVarP(&(sa.end_page),"end_page","e",-1,"specify page of end. default to -1.")
    flag.IntVarP(&(sa.page_len),"page_len","l",72,"specify length of page, default to 72.")
    flag.BoolVarP(&(sa.page_type),"page_type","f",false,"specify type of page, default to false.")
    flag.StringVarP(&(sa.print_dest),"print_dest","d","","specify print_dest if needed.")
    flag.Parse()

    targs:=flag.Args();
    if len(targs)>0 {
    	sa.in_filename=string(targs[0]);
    } else{
    	sa.in_filename="";
    }
    check(sa);

    if(sa.in_filename!=""){
	     _, errtest := os.Stat(sa.in_filename)
	    if os.IsNotExist(errtest) {
	        fmt.Fprintf(os.Stderr, "Open File Error\n")
	        os.Exit(1)
	    }
	} 

    var myin *os.File
    if sa.in_filename==""{
    	myin=os.Stdin;
    } else{
    	myin, _ = os.Open(sa.in_filename)
    }

    var myout io.WriteCloser
    var err error
	if len(sa.print_dest) == 0 {
		myout = os.Stdout
	} else {
		//本段代码借鉴了网上，因为我不了解管道的相关API
		cmd := exec.Command("lp", "-d"+sa.print_dest)
		cmd.Stdout, err = os.OpenFile(sa.print_dest, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		myout, err = cmd.StdinPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "-d pipe error! \n")
			os.Exit(1)
		}
		cmd.Run()
	}

	lineCounter:=0
	pageCounter:=1
	var mybuf string
	buf := bufio.NewReader(myin)
	if sa.page_type==false{
		for true{
				mybuf, err = buf.ReadString('\n')
	            lineCounter++
	            if lineCounter > sa.page_len {
	                pageCounter++
	                lineCounter = 1
	            }
	            if (pageCounter < sa.start_page) {
					continue;
				}
				if (pageCounter > sa.end_page){
					break;
				}
				if err!=nil{
					break
				}
				_, myerror := myout.Write([]byte(mybuf))
				if myerror != nil {
					break
				}
		}
	} else if sa.page_type==true{
		for true{
			mybuf, err = buf.ReadString('\f')
            pageCounter++
            if (pageCounter < sa.start_page) {
				continue;
			}
			if (pageCounter > sa.end_page){
				break;
			}
			if err!=nil{	
				break
			}
            _, myerror:=myout.Write([]byte(mybuf))
            if myerror != nil{
            	break
            }
		}
	}

	// if(pageCounter){
	// 	fmt.Fprintf(os.Stderr, "page number error! %d %d %d \n",pageCounter,sa.end_page,sa.start_page)
	// 	os.Exit(1)
	// }


}