<html>
    <head>
        <style>
            table {
                margin: auto;
                width: 40%;
                border: 2px solid black;
            }
            td {
                border: 1px solid lightgray;
            }
            .tr-even {
                background-color: white
            }
            .tr-odd {
                background-color: lightblue
            }
            tr {
                text-align: center;
            }
            .columnheader {
                color: white;
                background-color: green;
            }
            .emptylist {
                margin: auto;
                width: 300px;
                font-weight: bold;
                border: 3px solid lightblue;
                background-color: lightyellow;
                padding: 10px;
                text-align: center;
            }
        </style>
    </head>

    <body>
        {{if .}}
            <table>
                <tr class="columnheader">
                    <td><b>First Name</b></td>
                    <td><b>Last Name</b></td>
                </tr>
                {{$trclass := "tr-odd"}} 
                {{range $person := . }}
                    <tr class='{{$trclass}}'>
                        <td>{{$person.FirstName}}</td>
                        <td>{{$person.LastName}}</td>
                    </tr>
                    {{if (eq $trclass "tr-even")}}
                        {{$trclass = "tr-odd"}}
                    {{else}}
                        {{$trclass = "tr-even"}}
                    {{end}}
                {{end}}
            </table>
        {{else}}
            <div class="emptylist">Invalid Path</div>
        {{end}}
    </body>
</html>