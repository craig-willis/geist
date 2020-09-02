{{ prefix "prov" "http://www.w3.org/ns/prov#" }}
{{ prefix "provone" "http://purl.dataone.org/provone/2015/01/15/ontology#" }}
{{ prefix "wt" "http://wholetale.org/ontology/wt#" }}

{{ query "SelectRunID" '''
    SELECT ?r 
    WHERE {
        ?r a wt:TaleRun
    }
''' }}

{{ query "SelectTaleName" '''
    SELECT ?n 
    WHERE {
        <{{.}}> wt:TaleName ?n
    }
''' }}

{{ query "SelectRun" '''
    SELECT ?runID ?taleName ?runScript
    WHERE {
        ?runID a wt:TaleRun .
        ?runID wt:TaleName ?taleName .
        ?runID wt:TaleRunScript ?runScript
    }
''' }}

{{ macro "wt_node_run" "Label" '''
    "{{$Label}}"
''' }}


{{ macro "wt_run_node_style" '''
    node[shape=box style="filled" fillcolor="#FFFFFF" peripheries=2 fontname=Courier]
''' }}

{{ macro "wt_node_style_file" '''
    
    node[shape=box style="rounded,filled" fillcolor="#FFFFCC" peripheries=1 fontname=Helvetica]

''' }}

{{ query "SelectTaleOutputFiles" '''
    SELECT DISTINCT ?fileID ?filePath
    WHERE {
        ?e wt:ExecutionOf <{{.}}> .            
        ?p wt:ChildProcessOf ?e .   
        ?p wt:WroteFile ?fileID .          
        FILTER NOT EXISTS {
            ?_ wt:ReadFile ?fileID . 
        }
        ?fileID wt:FilePath ?filePath .
    }
    ORDER BY ?filePath
''' }}
}}}

{{ query "SelectTaleInputFiles" '''
    SELECT DISTINCT ?f ?fp
    WHERE {
        ?e wt:ExecutionOf <{{.}}> .               
        ?p wt:ChildProcessOf ?e .   
        ?p wt:ReadFile ?f .          
        FILTER NOT EXISTS {
            ?_ wt:WroteFile ?f . 
        }
        ?f wt:FilePath ?fp .
    }
    ORDER BY ?fp
''' }}
}}}

{{ query "SelectTaleOutputEdges" '''
    SELECT DISTINCT ?taleID ?filePath
    WHERE {
        ?e wt:ExecutionOf <{{.}}> .            
        ?p wt:ChildProcessOf ?e .   
        ?p wt:WroteFile ?fileID .          
        FILTER NOT EXISTS {
            ?_ wt:ReadFile ?fileID . 
        }
        ?fileID wt:FilePath ?filePath .
    }
    ORDER BY ?filePath
''' }}
}}}