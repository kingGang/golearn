var dirs=["a/b/11","a/b/12","11/22/33/44","11/测试","12/新建文件"]

var GetDirDepth=function(arr,title){
    if (arr==undefined||arr.length==0){
        return -1
    }
    for (b=0;b<arr.length;b++){
        if (arr[b].title==title){
            return b
        }
    }
    return -1
}

var tree=[];
len=dirs.length;
for (i=0;i<len;i++){
    dir=dirs[i]
    list=dir.split("/")
    listlen=list.length
    node=tree
    for (j=0;j<listlen;j++){
        depth=GetDirDepth(node,list[j])
        if (depth==-1){
            d=node.push({title:list[j],child:[]})
            depth=GetDirDepth(node,list[j])
            node=node[depth].child
        }else{
            node=node[depth].child
        }
    }
}
console.log(JSON.stringify(tree))