package futil

import (
//
)

var Banner string = `
    @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@@@@@[[\@@@@@@@@@@@@@@@@/[[@@@@@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@@/       ,@@@@@@@@@@        \@@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@@  ,@@@@    ,\@@/    ,@@@@   @@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@^  @@@@@@@@       ,@@@@@@@@  =@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@  ,@@@@@@@@@@@@@@@@@@@@@@@@^  @@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@  =@@@@@@@@@@@@@@@@@@@@@@@@^  \@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@^  @@@@@@@@@@@@@@@@@@@@@@@@@@  =@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@  =@@@@@@@@@@@@@@@@@@@@@@@@@@^  @@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@   @@@@@@@@@@@@@@@@@@@@@@@@@@@@  ,@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@   @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@   @@@@@@@@@@@@@
    @@@@@@@@@@@@   ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@   ,@@@@@@@@@@@
    @@@@@@@@@/   ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@    \@@@@@@@@
    @@@@@/     /@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\    ,\@@@@
    @@@    ]@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\     @@
    @@@       [[[\@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@/[[[      ,@@
    @@@@@@@\]]                                            ]]/@@@@@@
    @@@@@@@@@@@^       @@@@\]]]]]]]]]]]]]]]]/@@@@       =@@@@@@@@@@
    @@@@@@@@@@@   /\] =@@@@@@@@@@@@@@@@@@@@@@@@@@^ ]/\   @@@@@@@@@@
    @@@@@@@@@@^  @@@@@@@@@    =@@@@@@@@@@^    @@@@@@@@@  =@@@@@@@@@
    @@@@@@@@@@   @@@@@@@@@    ,@@@@@@@@@@     @@@@@@@@@  ,@@@@@@@@@
    @@@@@@@@@@\  ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@   /@@@@@@@@@
    @@@@@@@@@@@@   ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@    /@@@@@@@@@@
    @@@@@@@@@@@@@@     @@@@@@@@@@@@@@@@@@@@@@@@@@    ,/@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@  =@@@@@@@@@@@@@@@@@@@@@@@@^  @@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@\  ,@@@@@@@@@@@@@@@@@@@@@@   /@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@@\   \@@@@@@@@@@@@@@@@@@/   /@@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@@@@\   [@@@@@@@@@@@@@@[   /@@@@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@@@@@@\    ,\@@@@@@/     /@@@@@@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@@@@@@@@@\           ,/@@@@@@@@@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
    @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

    OOOOOOOOOOOOOOOOOO/[.                      .[OOOOOOOOOOOOOOOOOO
    OOOOOOOOOOoOOO/.                                ,\OOOOOOOOOOOOO
    OOOOOOOOOOO[                                        \OOOOOOOOOO
    OOOOOOOO/                      ...                    ,OOOOOOOO
    OOOOOO/             ,OOOOOOOOOOOOOOOOOOOOO^              \OOOOO
    OOOO/              ,OOOOOOOOOOOOOOOOOOOOOOO^              ,OOOO
    OOO                =OOOOOOOOOOOOOOOOOOOOOOOO                =OO
    O/                 =OOOOOOOOOOOOOOOOOOOOOOOO^                ,O
    /                  OOOOOOOOOOOOOOOOOOOOOOOOO^
                     .=OOOOOOOOOOOOOOOOOOOOOOOOOO .
                ]OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO 
               OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO^
                [OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO/ 
                   ,[OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO/ 
                       =OOOOOOOOOOOOOOOOOOOOOOOO^
                        OOOOOOOOOOOOOOOOOOOOOOOO
                        =OOOOOOOOOOOOOOOOOOOOOO^
                         \OOOOOOOOOOOOOOOOOOOO/
                         .OOOOOOOOOOOOOOOOOOOO 
                        \OOOOOOOOOOOOOOOOOOOOOOO.
                         \OOOOOOOOOOOOOOOOOOOOO.
                       .]]OOOOOOOOOOOOOOOOOOOOO]].
                .]/OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO\]
            ,/OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO]
    O     /OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO    ,
    OO\ ,OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO ,O
    OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO
    OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO
    OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO
    OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO
    OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO
    OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO

      ___           ___           ___           ___           ___           ___     
     /\  \         /\  \         /\  \         /\__\         /\  \         /\  \    
    /::\  \       /::\  \       /::\  \       /::|  |       /::\  \       /::\  \   
   /:/\:\  \     /:/\:\  \     /:/\:\  \     /:|:|  |      /:/\:\  \     /:/\:\  \  
  /::\~\:\  \   /::\~\:\  \   /::\~\:\  \   /:/|:|__|__   /::\~\:\  \   /::\~\:\  \ 
 /:/\:\ \:\__\ /:/\:\ \:\__\ /:/\:\ \:\__\ /:/ |::::\__\ /:/\:\ \:\__\ /:/\:\ \:\__\
 \/__\:\ \/__/ \/__\:\/:/  / \/_|::\/:/  / \/__/~~/:/  / \:\~\:\ \/__/ \/_|::\/:/  /
      \:\__\        \::/  /     |:|::/  /        /:/  /   \:\ \:\__\      |:|::/  / 
       \/__/        /:/  /      |:|\/__/        /:/  /     \:\ \/__/      |:|\/__/  
                   /:/  /       |:|  |         /:/  /       \:\__\        |:|  |    
                   \/__/         \|__|         \/__/         \/__/         \|__|    

  _________         __         _______       ____    ____    _________    _______    
 |_   ___  |       /  \       |_   __ \     |_   \  /   _|  |_   ___  |  |_   __ \  
   | |_  \_|      / /\ \        | |__) |      |   \/   |      | |_  \_|    | |__) | 
   |  _|         / ____ \       |  __ /       | |\  /| |      |  _|  _     |  __ /  
  _| |_        _/ /    \ \_    _| |  \ \_    _| |_\/_| |_    _| |___/ |   _| |  \ \_
 |_____|      |____|  |____|  |____| |___|  |_____||_____|  |_________|  |____| |___|
`
