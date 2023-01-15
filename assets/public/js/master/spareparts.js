$(document).ready( function () {
    
   var table = $('#myTable').DataTable( {
      
      processing: true,
      serverSide: true,
      paging:true,
      scrollX: true,
      scrollY: true,
    
      ajax: "/dashboard/master/daftar-sparepart",
      columns: [
              { data: 'nama', name: 'nama', class:'text-center' },
              { data: 'ukuran', name: 'ukuran' },
              { data: 'tipe', name: 'tipe' },
              { data: 'qty', name: 'qty' },
              { data: 'unit', name: 'unit' },
              { data: 'keterangan', name: 'keterangan' },            
          ],
    });

    $('#tombolRefreshTable').click(function(){
      table.draw();
    });
    
    $('#tombolTambahSparepart').click(function(){
      $('#saveBtnTambahSparepart').addClass('show');
      $('#saveBtnUpdateSparepart').attr('hidden', 'true'); 
      $('#saveBtnTambahSparepart').removeAttr('hidden'); 
    });

    // Tambah
    $('#saveBtnTambahSparepart').click( function(e){

      e.preventDefault();
      $(this).html('Sending..');
  
      $.ajax({
        data: $('#formtambahsparepart').serialize(),
        url: "/dashboard/master/daftar-sparepart",
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
                    $('#saveBtnTambahSparepart').html('Tambah');
                  });
                }else{
                  $('#formtambahsparepart').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahSparepart').html('Tambah');
                  $('#tombolTambahSparepart').trigger('click');
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses tambah data !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                    $(this).remove(); 
                 });
                    
                }
        },
        
      });

    })


    //edit
    $('body').on('click', '.editSparepart', function () {
      $('#saveBtnTambahSparepart').attr('hidden', 'true'); 
      $('#saveBtnUpdateSparepart').removeAttr('hidden'); 
      $('#saveBtnUpdateSparepart').html('Update');




      var sparepart_id = $(this).data('id');
      $.get("/dashboard/master/daftar-sparepart" +'/' + sparepart_id +'/edit', function (data) {
        $('#tambahsparepart').addClass('show');

          $('#id_sparepart').val(data.id);
          $('#nama').val(data.nama);
          $('#ukuran').val(data.ukuran);
          $('#satuan').val(data.satuan);
          $('#tipe').val(data.tipe);
          $('#qty').val(data.qty);
          $('#unit').val(data.unit);
          $('#keterangan').val(data.keterangan);

          $("html, body").animate({ scrollTop: 0 }, "slow");
          // $('#scroll-ku').scrollTop = 0;
          
      })
   });


    //update
   $('#saveBtnUpdateSparepart').click(function (e) {

      e.preventDefault();
      $(this).html('Sending..');
      var id_sparepart = $('#id_sparepart').val();
      $.ajax({
        data: $('#formtambahsparepart').serialize(),
        url: "/dashboard/master/daftar-sparepart"+'/'+id_sparepart,
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
                    $('#updateBtnUpdateSparepart').html('Update');
                  });
                }else{
                  $(this).html('Update');
                  $('#formtambahsparepart').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahSparepart').html('Tambah');
                  $('#tombolTambahSparepart').trigger('click');   
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses Update !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                     $(this).remove(); 
                  });
                }
        },
        
      });

    })



    //batal
   $('#saveBtnBatalSparepart').click( function(e){
      $(this).html('Sending..');
      e.preventDefault();
      $('#formtambahSparepart').trigger("reset");
         $('#saveBtnBatalSparepart').html('Batal');
         $('#tombolTambahSparepart').trigger('click');
         $(document).find('span.error-text').text('');
 
    })



    //delete
    $('body').on('click', '.deleteSparepart', function () {
     
      var sparepart_id = $(this).data("id");
      var r = confirm("Are You sure want to delete !");
       if (r == true) {
         $.ajax({
          type: "DELETE",
          url: "/dashboard/master/daftar-sparepart"+'/'+sparepart_id,
          success: function (data) {
            $('.table-responsive').before(`<div class="alert alert-danger alert-dismissible fade show" role="alert"> Sukses hapus sparepart!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
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