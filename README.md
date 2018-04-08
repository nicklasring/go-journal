# go-journal
Dirty CLI-Journal - Experimenting with Burntsushi xlib bindings

### installation

go get github.com/nicklasring/go-journal


### using

```shell
# Creates a journal in $HOME/.journal/YYYY-MM-DD/journal-1.txt
./journal -t "Today i did this.. and that.." 

# Creates a journal in $HOME/.journal/YYYY-MM-DD/journal-2.txt
# Creates a screenshot in $HOME/.journal/YYYY-MM-DD/journal-2.png
./journal -i "Today i did that.. and this.."
```



### TODO

- [ ] Write better structure
- [x] Implement screenshotting feature
- [ ] Implement rendering of borders for selected area during click-release for screenshot (like imagemagick import)
- [ ] Implement Markdown formatting via "special syntax" eg. "h1:My Header:list:Item 1:Item 2" etc...
