#!/usr/bin/env bash

RUNNER='../../common/run_script_example.sh'

# *****************************************************************************

bash ${RUNNER} SETUP "IMPORT SDTL" << END_SCRIPT

blazegraph destroy --dataset kb
blazegraph create --dataset kb
blazegraph import --format jsonld --file ../data/compute-sdtl.jsonld

END_SCRIPT



# *****************************************************************************

bash ${RUNNER} Q1 "WHAT COMMANDS ARE EXECUTED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?line

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q2 "WHAT DATA FILES ARE LOADED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?file_name ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command rdf:type sdtl:Load .
        ?command sdtl:FileName ?file_name .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q3 "WHAT DATA FILES ARE SAVED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?file_name ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command rdf:type sdtl:Save .
        ?command sdtl:FileName ?file_name .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q4 "WHAT VARIABLES ARE LOADED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?loaded_variable ?dataframe ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command rdf:type sdtl:Load .
        ?command sdtl:ProducesDataframe ?dataframe_description .
        ?dataframe_description sdtl:DataframeName ?dataframe .
        ?dataframe_description sdtl:VariableInventory ?loaded_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q5 "WHAT VARIABLES ARE SAVED BY THE SCRIPT?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?saved_variable ?dataframe ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command rdf:type sdtl:Save .
        ?command sdtl:ProducesDataframe ?dataframe_description .
        ?dataframe_description sdtl:DataframeName ?dataframe .
        ?dataframe_description sdtl:VariableInventory ?saved_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    }

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q6 "WHAT COMMANDS UPDATE EACH DATAFRAME?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?dataframe ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command sdtl:ProducesDataframe ?dataframe_description .
        ?dataframe_description sdtl:DataframeName ?dataframe .
        ?dataframe_description sdtl:VariableInventory ?variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?dataframe ?source_line

__END_QUERY__

END_SCRIPT



bash ${RUNNER} Q7 "WHAT COMMANDS UPDATE EACH VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?updated_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command sdtl:Variable ?variable .
        ?variable sdtl:VariableName ?updated_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?variable ?source_line

__END_QUERY__

END_SCRIPT


bash ${RUNNER} Q8 "WHAT COMMANDS USE EACH VARIABLE?" << END_SCRIPT

blazegraph select --format table << __END_QUERY__

    PREFIX sdtl: <https://rdf-vocabulary.ddialliance.org/sdtl#>

    SELECT DISTINCT ?used_variable ?command ?source_line ?source_text
    WHERE {
        ?program rdf:type sdtl:Program .
        ?program sdtl:Commands ?command .
        ?command sdtl:Expression ?expression .
        ?expression (sdtl:Arguments/sdtl:ArgumentValue)+/sdtl:VariableName ?used_variable .
        ?command sdtl:SourceInformation ?source_info .
        ?source_info sdtl:LineNumberStart ?source_line .
        ?source_info sdtl:OriginalSourceText ?source_text .
    } ORDER BY ?used_variable ?source_line

__END_QUERY__

END_SCRIPT
