
{{ macro "gv_graph" "Name" '''
    digraph {{$Name}} { 
    rankdir=LR
''' }}

{{ macro "gv_title" "Title" '''
    fontname=Courier; fontsize=18; labelloc=t
    label="{{$Title}}"
''' }}

{{ macro "gv_end" '''
    }
''' }}

{{ macro "gv_cluster" "Name" '''
    subgraph {{ printf "cluster_%s" $Name }} { label=""; color=white; penwidth=0
    subgraph {{ printf "cluster_%s_inner" $Name }} { label=""; color=white
''' }}

{{ macro "gv_cluster_end" '''
    }}
''' }}

{{ macro "labeled_node" "NodeID" "NodeLabel" '''
    "{{$NodeID}}" [label="{{$NodeLabel}}"]
''' }}

{{ macro "gv_edge" "Head" '''
    "08-branched-pipeline" -> "{{$Head}}"
''' }}

{{ macro "gv_input_edge" "Tail" '''
    "{{$Tail}}" -> "08-branched-pipeline"
''' }}