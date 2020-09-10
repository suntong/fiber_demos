/*
   Licensed to the Apache Software Foundation (ASF) under one or more
   contributor license agreements.  See the NOTICE file distributed with
   this work for additional information regarding copyright ownership.
   The ASF licenses this file to You under the Apache License, Version 2.0
   (the "License"); you may not use this file except in compliance with
   the License.  You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
var showControllersOnly = false;
var seriesFilter = "^(T\\d+_.*)?$";
var filtersOnlySampleSeries = true;

/*
 * Add header in statistics table to group metrics by category
 * format
 *
 */
function summaryTableHeader(header) {
    var newRow = header.insertRow(-1);
    newRow.className = "tablesorter-no-sort";
    var cell = document.createElement('th');
    cell.setAttribute("data-sorter", false);
    cell.colSpan = 1;
    cell.innerHTML = "Requests";
    newRow.appendChild(cell);

    cell = document.createElement('th');
    cell.setAttribute("data-sorter", false);
    cell.colSpan = 3;
    cell.innerHTML = "Executions";
    newRow.appendChild(cell);

    cell = document.createElement('th');
    cell.setAttribute("data-sorter", false);
    cell.colSpan = 6;
    cell.innerHTML = "Response Times (ms)";
    newRow.appendChild(cell);

    cell = document.createElement('th');
    cell.setAttribute("data-sorter", false);
    cell.colSpan = 1;
    cell.innerHTML = "Throughput";
    newRow.appendChild(cell);

    cell = document.createElement('th');
    cell.setAttribute("data-sorter", false);
    cell.colSpan = 2;
    cell.innerHTML = "Network (KB/sec)";
    newRow.appendChild(cell);
}

/*
 * Populates the table identified by id parameter with the specified data and
 * format
 *
 */
function createTable(table, info, formatter, defaultSorts, seriesIndex, headerCreator) {
    var tableRef = table[0];

    // Create header and populate it with data.titles array
    var header = tableRef.createTHead();

    // Call callback is available
    if(headerCreator) {
        headerCreator(header);
    }

    var newRow = header.insertRow(-1);
    for (var index = 0; index < info.titles.length; index++) {
        var cell = document.createElement('th');
        cell.innerHTML = info.titles[index];
        newRow.appendChild(cell);
    }

    var tBody;

    // Create overall body if defined
    if(info.overall){
        tBody = document.createElement('tbody');
        tBody.className = "tablesorter-no-sort";
        tableRef.appendChild(tBody);
        var newRow = tBody.insertRow(-1);
        var data = info.overall.data;
        for(var index=0;index < data.length; index++){
            var cell = newRow.insertCell(-1);
            cell.innerHTML = formatter ? formatter(index, data[index]): data[index];
        }
    }

    // Create regular body
    tBody = document.createElement('tbody');
    tableRef.appendChild(tBody);

    var regexp;
    if(seriesFilter) {
        regexp = new RegExp(seriesFilter, 'i');
    }
    // Populate body with data.items array
    for(var index=0; index < info.items.length; index++){
        var item = info.items[index];
        if((!regexp || filtersOnlySampleSeries && !info.supportsControllersDiscrimination || regexp.test(item.data[seriesIndex]))
                &&
                (!showControllersOnly || !info.supportsControllersDiscrimination || item.isController)){
            if(item.data.length > 0) {
                var newRow = tBody.insertRow(-1);
                for(var col=0; col < item.data.length; col++){
                    var cell = newRow.insertCell(-1);
                    cell.innerHTML = formatter ? formatter(col, item.data[col]) : item.data[col];
                }
            }
        }
    }

    // Add support of columns sort
    table.tablesorter({sortList : defaultSorts});
}

$(document).ready(function() {

    // Customize table sorter default options
    $.extend( $.tablesorter.defaults, {
        theme: 'blue',
        cssInfoBlock: "tablesorter-no-sort",
        widthFixed: true,
        widgets: ['zebra']
    });

    var data = {"OkPercent": 97.36842105263158, "KoPercent": 2.6315789473684212};
    var dataset = [
        {
            "label" : "KO",
            "data" : data.KoPercent,
            "color" : "#FF6347"
        },
        {
            "label" : "OK",
            "data" : data.OkPercent,
            "color" : "#9ACD32"
        }];
    $.plot($("#flot-requests-summary"), dataset, {
        series : {
            pie : {
                show : true,
                radius : 1,
                label : {
                    show : true,
                    radius : 3 / 4,
                    formatter : function(label, series) {
                        return '<div style="font-size:8pt;text-align:center;padding:2px;color:white;">'
                            + label
                            + '<br/>'
                            + Math.round10(series.percent, -2)
                            + '%</div>';
                    },
                    background : {
                        opacity : 0.5,
                        color : '#000'
                    }
                }
            }
        },
        legend : {
            show : true
        }
    });

    // Creates APDEX table
    createTable($("#apdexTable"), {"supportsControllersDiscrimination": true, "overall": {"data": [0.9528301886792453, 500, 1500, "Total"], "isController": false}, "titles": ["Apdex", "T (Toleration threshold)", "F (Frustration threshold)", "Label"], "items": [{"data": [1.0, 500, 1500, "createOrder"], "isController": false}, {"data": [1.0, 500, 1500, "T02_ Login page"], "isController": true}, {"data": [1.0, 500, 1500, "T03_ Login submit"], "isController": true}, {"data": [1.0, 500, 1500, "T05_ SignOff"], "isController": true}, {"data": [1.0, 500, 1500, "T01_ Home page"], "isController": true}, {"data": [1.0, 500, 1500, "placeOrder"], "isController": false}, {"data": [1.0, 500, 1500, "viewAccount"], "isController": false}, {"data": [1.0, 500, 1500, "signOff"], "isController": false}, {"data": [1.0, 500, 1500, "signonForm"], "isController": false}, {"data": [1.0, 500, 1500, "viewProduct"], "isController": false}, {"data": [1.0, 500, 1500, "confirmOrder"], "isController": false}, {"data": [0.0, 500, 1500, "JDBC Request"], "isController": false}, {"data": [1.0, 500, 1500, "viewCategory"], "isController": false}, {"data": [0.5, 500, 1500, "T04_ Create Order"], "isController": true}, {"data": [1.0, 500, 1500, "viewCatalog"], "isController": false}, {"data": [1.0, 500, 1500, "viewCart"], "isController": false}, {"data": [1.0, 500, 1500, "OS Process Sampler"], "isController": false}]}, function(index, item){
        switch(index){
            case 0:
                item = item.toFixed(3);
                break;
            case 1:
            case 2:
                item = formatDuration(item);
                break;
        }
        return item;
    }, [[0, 0]], 3);

    // Create statistics table
    createTable($("#statisticsTable"), {"supportsControllersDiscrimination": true, "overall": {"data": ["Total", 38, 1, 2.6315789473684212, 139.5526315789474, 106, 711, 165.20000000000024, 375.649999999999, 711.0, 7.042253521126761, 2.710521972294292, 6.211741046608599], "isController": false}, "titles": ["Label", "#Samples", "KO", "Error %", "Average", "Min", "Max", "90th pct", "95th pct", "99th pct", "Transactions\/s", "Received", "Sent"], "items": [{"data": ["createOrder", 3, 0, 0.0, 107.33333333333333, 107, 108, 108.0, 108.0, 108.0, 1.0578279266572639, 0.42561045486600846, 0.9555574532792667], "isController": false}, {"data": ["T02_ Login page", 3, 0, 0.0, 106.66666666666667, 106, 107, 107.0, 107.0, 107.0, 1.053370786516854, 0.42381715238764045, 0.9186133909761236], "isController": true}, {"data": ["T03_ Login submit", 3, 0, 0.0, 221.66666666666666, 219, 226, 226.0, 226.0, 226.0, 1.0138560324433932, 0.7940552129097668, 2.010880470598175], "isController": true}, {"data": ["T05_ SignOff", 3, 0, 0.0, 214.0, 213, 215, 215.0, 215.0, 215.0, 1.021450459652707, 0.8079832737487233, 1.8084860188968337], "isController": true}, {"data": ["T01_ Home page", 3, 0, 0.0, 204.66666666666666, 146, 320, 320.0, 320.0, 320.0, 0.9430996541967935, 0.36931929817667397, 0.5461504833385727], "isController": true}, {"data": ["placeOrder", 3, 0, 0.0, 127.0, 116, 135, 135.0, 135.0, 135.0, 1.0496850944716585, 0.4090081569279216, 1.5150728219034288], "isController": false}, {"data": ["viewAccount", 3, 0, 0.0, 114.33333333333333, 112, 119, 119.0, 119.0, 119.0, 1.0518934081346423, 0.41192310220897616, 1.1412632582398317], "isController": false}, {"data": ["signOff", 3, 0, 0.0, 107.0, 107, 107, 107.0, 107.0, 107.0, 1.0596962204168139, 0.42325757241257506, 0.9448268058989756], "isController": false}, {"data": ["signonForm", 3, 0, 0.0, 106.66666666666667, 106, 107, 107.0, 107.0, 107.0, 1.053001053001053, 0.4236683924183924, 0.9182909573534573], "isController": false}, {"data": ["viewProduct", 3, 0, 0.0, 107.66666666666667, 107, 109, 109.0, 109.0, 109.0, 1.056338028169014, 0.4456426056338028, 0.9727800396126761], "isController": false}, {"data": ["confirmOrder", 3, 0, 0.0, 107.66666666666667, 107, 108, 108.0, 108.0, 108.0, 1.0596962204168139, 0.4377456066760862, 0.9334433504062168], "isController": false}, {"data": ["JDBC Request", 1, 1, 100.0, 711.0, 711, 711, 711.0, 711.0, 711.0, 1.4064697609001406, 0.09614539381153306, 0.0], "isController": false}, {"data": ["viewCategory", 3, 0, 0.0, 107.33333333333333, 107, 108, 108.0, 108.0, 108.0, 1.055594651653765, 0.44739070197044334, 0.9432315490851513], "isController": false}, {"data": ["T04_ Create Order", 3, 0, 0.0, 664.3333333333334, 653, 674, 674.0, 674.0, 674.0, 0.8823529411764707, 2.1843405330882355, 5.265682444852941], "isController": true}, {"data": ["viewCatalog", 9, 0, 0.0, 139.66666666666666, 106, 320, 320.0, 320.0, 320.0, 2.0847810979847115, 0.8164035354412786, 1.6375575486448921], "isController": false}, {"data": ["viewCart", 3, 0, 0.0, 107.33333333333333, 107, 108, 108.0, 108.0, 108.0, 1.0570824524312896, 0.4490535808668076, 0.9786271141649049], "isController": false}, {"data": ["OS Process Sampler", 1, 0, 0.0, 358.0, 358, 358, 358.0, 358.0, 358.0, 2.793296089385475, 0.07365136173184358, 0.0], "isController": false}]}, function(index, item){
        switch(index){
            // Errors pct
            case 3:
                item = item.toFixed(2) + '%';
                break;
            // Mean
            case 4:
            // Mean
            case 7:
            // Percentile 1
            case 8:
            // Percentile 2
            case 9:
            // Percentile 3
            case 10:
            // Throughput
            case 11:
            // Kbytes/s
            case 12:
            // Sent Kbytes/s
                item = item.toFixed(2);
                break;
        }
        return item;
    }, [[0, 0]], 0, summaryTableHeader);

    // Create error table
    createTable($("#errorsTable"), {"supportsControllersDiscrimination": false, "titles": ["Type of error", "Number of errors", "% in errors", "% in all samples"], "items": [{"data": ["S0001 4832\/com.microsoft.sqlserver.jdbc.SQLServerException: Bulk load: An unexpected end of file was encountered in the data file.", 1, 100.0, 2.6315789473684212], "isController": false}]}, function(index, item){
        switch(index){
            case 2:
            case 3:
                item = item.toFixed(2) + '%';
                break;
        }
        return item;
    }, [[1, 1]]);

        // Create top5 errors by sampler
    createTable($("#top5ErrorsBySamplerTable"), {"supportsControllersDiscrimination": false, "overall": {"data": ["Total", 38, 1, "S0001 4832\/com.microsoft.sqlserver.jdbc.SQLServerException: Bulk load: An unexpected end of file was encountered in the data file.", 1, null, null, null, null, null, null, null, null], "isController": false}, "titles": ["Sample", "#Samples", "#Errors", "Error", "#Errors", "Error", "#Errors", "Error", "#Errors", "Error", "#Errors", "Error", "#Errors"], "items": [{"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": ["JDBC Request", 1, 1, "S0001 4832\/com.microsoft.sqlserver.jdbc.SQLServerException: Bulk load: An unexpected end of file was encountered in the data file.", 1, null, null, null, null, null, null, null, null], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}, {"data": [], "isController": false}]}, function(index, item){
        return item;
    }, [[0, 0]], 0);

});
