# Git Cmd
'''
操作不熟悉 git cmd, 並記錄期操作步驟狀態以及結果
'''

* Base operte > Git mod init, import other file 
    * step 1
        * go mod init github.com/hsineternity/GitCmd 

* Git cmd operate
    * rebase : At A branch input cmd > git rebase B
        * describe : new commit of A branch is added to branch B
        * Ans : 
            因為有 merge 衝突, 所以 add cat print 內容沒有直接加上去
            <!-- A branch git log  -->
            1c2a68a (HEAD -> cat, origin/cat) add cat print
            781d9a2 add cat func
            06d21ec (origin/master, master) fix main content
            <!-- B branch git log  -->
            9d9a99a (HEAD -> dog, origin/dog) add dog print
            d0d1f87 add dog func
            06d21ec (origin/master, master) fix main content
            <!-- A branch B git log  -->
            8b05f32 (HEAD) add cat func
            9d9a99a (origin/dog, dog) add dog print
            d0d1f87 add dog func
            06d21ec (origin/master, master) fix main content
        * questrion :
            git push origin HEAD:<name-of-remote-branch>
            rebase 完後會遇到需要指定合回去的 HEAD 狀況, 這邊因為是 cat rebase dog,
            log 是 cat 附加在 dog 上, 故將 push HEAD 指向 dog
