$(function () {

  $('#selectsparepart').select2();

    
  $.ajaxSetup({
    headers: {
        'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
    }
  });

  var table= $('#myTable').DataTable( {
      
    processing: true,
    serverSide: true,
    paging:true,
    scrollX: true,
    scrollY: true,
  
    ajax: "/dashboard/form/penambahan-stock-sparepart",
      columns: [
              { data: 'sparepart.nama', name: 'sparepart_nama', orderable:false },
              { data: 'user.nama_lengkap', name: 'nama_lengkap', orderable:false },
              { data: 'qty', name: 'qty', orderable:false },
              { data: 'no_invoice', name: 'no_invoice', orderable:false },
              { data: 'created_at', name: 'created_at', orderable:false },
              { data: 'action', name: 'action', orderable:false },
          ],
    });

    $('#tombolRefreshTable').click(function(){
      table.draw();
    });
    
    $('#tombolTambahStockSparepart').click(function(){
      $('#keyword').focus(); 
      $('#saveBtnTambahStockSparepart').addClass('show');
      $('#saveBtnUpdateStockSparepart').attr('hidden', 'true'); 
      $('#saveBtnTambahStockSparepart').removeAttr('hidden'); 
      
    });
   

    // Tambah
    $('#saveBtnTambahStockSparepart').click( function(e){

      e.preventDefault();
      $(this).html('Sending..');
  
      $.ajax({
        data: $('#formtambahstocksparepart').serialize(),
        url: "/dashboard/form/penambahan-stock-sparepart",
        type: "POST",
        dataType: 'json',
        beforeSend:function(){
              $(document).find('span.error-text').text('');
            },
        success: function (data) {

          
                if(data.status==0){
                  $.each(data.error, function(frefix, val){
                    $('span.'+frefix+'_error').text(val[0]);
                    $('span.'+frefix+'_errornotif').text('['+val[0]+']');
                    $('#saveBtnTambahStockSparepart').html('Tambah');
                  });
                }else{
                  $('#formtambahstocksparepart').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahStockSparepart').html('Tambah');
                  $('#tombolTambahStockSparepart').trigger('click');
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses tambah data !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                    $(this).remove(); 
                 });
                    
                }
        },
        
      });

    })

   

    //edit
    $('body').on('click', '.editStockSparepart', function () {
      $('#saveBtnTambahStockSparepart').attr('hidden', 'true'); 
      $('#saveBtnUpdateStockSparepart').removeAttr('hidden'); 
      $('#saveBtnUpdateStockSparepart').html('Update');




      var stock_sparepart_id = $(this).data('id');
      $.get("/dashboard/form/penambahan-stok-sparepart" +'/' + stock_sparepart_id +'/edit', function (data) {
        $('#tambahstocksparepart').addClass('show');

          $('#selectsparepart').val(data.id);
          $('#qty').val(data.qty);
          $('#no_invoice').val(data.no_invoice);

          $("html, body").animate({ scrollTop: 0 }, "slow");
          // $('#scroll-ku').scrollTop = 0;
          
      })
   });


    //update
    $('#saveBtnUpdateStockSparepart').click(function (e) {

      e.preventDefault();
      $(this).html('Sending..');
      var id_stock_sparepart = $('#id_stock_sparepart').val();
      $.ajax({
        data: $('#formtambahmesin').serialize(),
        url: "/dashboard/form/penambahan-stock-sparepart"+'/'+id_stock_sparepart,
        type: "PUT",
        dataType: 'json',
        beforeSend:function(){
              $(document).find('span.error-text').text('');
            },
        success: function (data) {

          
                if(data.status==0){
                  $.each(data.error, function(frefix, val){
                    $('span.'+frefix+'_error').text(val[0]);
                    $('span.'+frefix+'_errornotif').text('['+val[0]+']');
                    $('#updateBtnUpdateStockSparepart').html('Update');
                  });
                }else{
                  $(this).html('Update');
                  $('#formtambahstocksparepart').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahStockSparepart').html('Tambah');
                  $('#tombolTambahStockSparepart').trigger('click');   
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses Update !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                     $(this).remove(); 
                  });
                }
        },
        
      });

    })


    //batal
   $('#saveBtnBatalStockSparepart').click( function(e){
     $(this).html('Sending..');
     e.preventDefault();
     $('#formtambahstocksparepart').trigger("reset");
        $('#saveBtnBatalStockSparepart').html('Batal');
        $('#tombolTambahStockSparepart').trigger('click');
        $(document).find('span.error-text').text('');

   })


    //delete
    $('body').on('click', '.deleteStockSparepart', function () {
     
        var stock_sparepart_id = $(this).data("id");
        var r = confirm("Are You sure want to delete !");
         if (r == true) {
           $.ajax({
            type: "DELETE",
            url: "/dashboard/form/penambahan-stok-sparepart"+'/'+stock_sparepart_id,
            success: function (data) {
              $('.table-responsive').before(`<div class="alert alert-danger alert-dismissible fade show" role="alert"> Sukses hapus stock sparepart!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
              table.draw();
              $('.alert').delay(2000).fadeOut(function() {
                $(this).remove(); 
             });
            },
            error: function (data) {
                console.log('Error:', data);
            }
           });
         } else {
           table.draw();
         }
    });

  

  } );