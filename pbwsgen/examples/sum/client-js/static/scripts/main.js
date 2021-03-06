requirejs.config({
    baseUrl: 'scripts/lib',

    paths: {
        'ByteBuffer': 'ByteBufferAB.min',
        'ProtoBuf': 'ProtoBuf.min',
        'Long': 'Long.min',
        'jquery': 'jquery-2.1.3.min'
    },

    shim: {
        'jquery.binarytransport': {
            deps: ['jquery'],
            exports: 'BinaryTransport'
        }
    }
});

requirejs(['jquery', 'sum.client'],
    function($, sum) {
        $("#calculateSum").click(function(ev) {
            // Create a req message
            // Note: any value errors are caught here
            // TODO: catch any errors and do something useful
            var sumReq = sum.createReq({
                "a": parseInt($("#valueA").val()),
                "b": parseInt($("#valueB").val())
            })

            // Call web service with callback function
            // TODO: handle errors
            // TODO: promises etc? Whats the new stuff in js land?
            sum.call("http://localhost:8080/sum/ws", sumReq, function(sumResp) {
                $("#valueResult").html(sumResp.sum)
            });
        });
    });