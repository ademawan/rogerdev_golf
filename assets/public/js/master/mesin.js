$(function () {
  
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
    
      ajax: "/dashboard/master/daftar-mesin",
      columns: [
              { data: 'jenis_mesin', name: 'jenis_mesin', class:'text-center' },
              { data: 'brand', name: 'brand' },
              { data: 'negara_asal', name: 'negara_asal', orderable:false },
              { data: 'qty', name: 'qty', orderable:false },
              { data: 'tahun', name: 'tahun', orderable:false },
              { data: 'no_registrasi', name: 'no_registrasi', orderable:false },
              { data: 'daya_kilowatt', name: 'daya_kilowatt', orderable:false },
              { data: 'daya_voltase', name: 'daya_voltase', orderable:false },
              { data: 'kondisi', name: 'kondisi', orderable:false },
              { data: 'lokasi_mesin', name: 'lokasi_mesin', orderable:false },
              { data: 'keterangan', name: 'keterangan', orderable:false },
              { data: 'action', name: 'action', orderable:false }
          ]
    });
    

    $('#tombolRefreshTable').click(function(){
      table.draw();
    });
    
    $('#tombolTambahMesin').click(function(){
      $('#saveBtnTambahMesin').addClass('show');
      $('#saveBtnUpdateMesin').attr('hidden', 'true'); 
      $('#saveBtnTambahMesin').removeAttr('hidden'); 
    });


    // Tambah mesin
    $('#saveBtnTambahMesin').click( function(e){

      e.preventDefault();
      $(this).html('Sending..');
  
      $.ajax({
        data: $('#formtambahmesin').serialize(),
        url: "/dashboard/master/daftar-mesin",
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
                    $('#saveBtnTambahMesin').html('Tambah');
                  });
                }else{
                  $('#formtambahmesin').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahMesin').html('Tambah');
                  $('#tombolTambahMesin').trigger('click');
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses tambah data !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                    $(this).remove(); 
                 });
                    
                }
        },
        
      });

    })

   

    //edit
    $('body').on('click', '.editMesin', function () {
      $('#saveBtnTambahMesin').attr('hidden', 'true'); 
      $('#saveBtnUpdateMesin').removeAttr('hidden'); 
      $('#saveBtnUpdateMesin').html('Update');




      var mesin_id = $(this).data('id');
      $.get("/dashboard/master/daftar-mesin" +'/' + mesin_id +'/edit', function (data) {
        $('#tambahmesin').addClass('show');

          $('#id_mesin').val(data.id);
          $('#jenis_mesin').val(data.jenis_mesin);
          $('#brand').val(data.brand);
          $('#negara_asal').val(data.negara_asal);
          $('#qty').val(data.qty);
          $('#tahun').val(data.tahun);
          $('#no_registrasi').val(data.no_registrasi);
          $('#daya_kilowatt').val(data.daya_kilowatt);
          $('#daya_voltase').val(data.daya_voltase);
          $('#kondisi').val(data.kondisi);
          $('#lokasi_mesin').val(data.lokasi_mesin);
          $('#keterangan').val(data.keterangan);
          $('#status').val(data.status);

          $("html, body").animate({ scrollTop: 0 }, "slow");
          // $('#scroll-ku').scrollTop = 0;
          
      })
   });


    //update
    $('#saveBtnUpdateMesin').click(function (e) {

      e.preventDefault();
      $(this).html('Sending..');
      var id_mesin = $('#id_mesin').val();
      $.ajax({
        data: $('#formtambahmesin').serialize(),
        url: "/dashboard/master/daftar-mesin"+'/'+id_mesin,
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
                    $('#updateBtnUpdateMesin').html('Update');
                  });
                }else{
                  $(this).html('Update');
                  $('#formtambahmesin').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahMesin').html('Tambah');
                  $('#tombolTambahMesin').trigger('click');   
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses Update !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                     $(this).remove(); 
                  });
                }
        },
        
      });

    })


    //batal
   $('#saveBtnBatalMesin').click( function(e){
     $(this).html('Sending..');
     e.preventDefault();
     $('#formtambahmesin').trigger("reset");
        $('#saveBtnBatalMesin').html('Batal');
        $('#tombolTambahMesin').trigger('click');
        $(document).find('span.error-text').text('');

   })


    //delete
    $('body').on('click', '.deleteMesin', function () {
     
        var mesin_id = $(this).data("id");
        var r = confirm("Are You sure want to delete !");
         if (r == true) {
           $.ajax({
            type: "DELETE",
            url: "/dashboard/master/daftar-mesin"+'/'+mesin_id,
            success: function (data) {
              $('.table-responsive').before(`<div class="alert alert-danger alert-dismissible fade show" role="alert"> Sukses hapus mesin!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
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