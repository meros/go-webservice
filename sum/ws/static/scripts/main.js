requirejs.config({
    //By default load any module IDs from js/lib
    baseUrl: 'scripts/lib',
    //except, if the module ID starts with "app",
    //load it from the js/app directory. paths
    //config is relative to the baseUrl, and
    //never includes a ".js" extension since
    //the paths config could be for a directory.
    paths: {
        'ByteBuffer': 'ByteBufferAB.min',
	'ProtoBuf': 'ProtoBuf.min',
	'Long': 'Long.min',
	'jquery': 'jquery-2.1.3.min'
    },

    shim: {
	'jquery.binarytransport': {
            //These script dependencies should be loaded before loading
            //backbone.js
            deps: ['jquery'],
            //Once loaded, use the global 'Backbone' as the
            //module value.
            exports: 'BinaryTransport'
	}
    }
});

// Start the main app logic.
requirejs(['jquery', 'sum.pbwsclient'],
	  function   ($, sum) {
              // If a user clicks on it, say hello!
              $("#calculateSum").click(function(ev) {
                  var sumReq = sum.createReq({
                      "a": parseInt($("#valueA").val()),
                      "b": parseInt($("#valueB").val())
                  })

                  sum.call(sumReq, function(sumResp) {
		      $("#valueResult").html(sumResp.sum)
		  });                
              });
          });
