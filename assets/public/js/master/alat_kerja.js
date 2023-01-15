$(document).ready( function () {
    
   var table = $('#myTable').DataTable( {
      
      processing: true,
      serverSide: true,
      paging:true,
      scrollX: true,
      scrollY: true,
    
      ajax: "http://telkomsel.localhost:8088/admin/user/datatables",
      columns: [
              { data: 'nama', name: 'nama', class:'text-center' },
              { data: 'ukuran', name: 'ukuran' },
              { data: 'merk', name: 'merk' },
              { data: 'qty', name: 'qty' },
              { data: 'kondisi', name: 'kondisi' },
              { data: 'keterangan', name: 'keterangan' },            
              { data: 'no_registrasi', name: 'no_registrasi' },            
              { data: 'action', name: 'action' },            
          ],
    });

    $('#tombolRefreshTable').click(function(){
      table.draw();
    });
    
    $('#tombolTambahAlatKerja').click(function(){
      $('#saveBtnTambahAlatKerja').addClass('show');
      $('#saveBtnUpdateAlatKerja').attr('hidden', 'true'); 
      $('#saveBtnTambahAlatKerja').removeAttr('hidden'); 
    });

    // Tambah alat kerja
    $('#saveBtnTambahAlatKerja').click( function(e){

      e.preventDefault();
      $(this).html('Sending..');
  
      $.ajax({
        data: $('#formtambahalatkerja').serialize(),
        url: "/dashboard/master/daftar-alat-kerja",
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
                    $('#saveBtnTambahAlatKerja').html('Tambah');
                  });
                }else{
                  $('#formtambahalatkerja').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahAlatKerja').html('Tambah');
                  $('#tombolTambahAlatKerja').trigger('click');
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses tambah data !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                    $(this).remove(); 
                 });
                    
                }
        },
        
      });

    })


    //edit
    $('body').on('click', '.editAlatKerja', function () {
      $('#saveBtnTambahAlatKerja').attr('hidden', 'true'); 
      $('#saveBtnUpdateAlatKerja').removeAttr('hidden'); 
      $('#saveBtnUpdateAlatKerja').html('Update');




      var alat_kerja_id = $(this).data('id');
      $.get("/dashboard/master/daftar-alat-kerja" +'/' + alat_kerja_id +'/edit', function (data) {
        $('#tambahalatkerja').addClass('show');

          $('#id_alatkerja').val(data.id);
          $('#nama').val(data.nama);
          $('#ukuran').val(data.ukuran);
          $('#satuan').val(data.satuan_id);
          $('#merk').val(data.merk);
          $('#qty').val(data.qty);
          $('#unit').val(data.unit_id);
          $('#kondisi').val(data.kondisi);
          $('#keterangan').val(data.keterangan);
          $('#no_registrasi').val(data.no_registrasi);

          $("html, body").animate({ scrollTop: 0 }, "slow");
          // $('#scroll-ku').scrollTop = 0;
          
      })
   });


    //update
   $('#saveBtnUpdateAlatKerja').click(function (e) {

      e.preventDefault();
      $(this).html('Sending..');
      var id_alatkerja = $('#id_alatkerja').val();
      $.ajax({
        data: $('#formtambahalatkerja').serialize(),
        url: "/dashboard/master/daftar-alatkerja"+'/'+id_alatkerja,
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
                    $('#updateBtnUpdateAlatKerja').html('Update');
                  });
                }else{
                  $(this).html('Update');
                  $('#formtambahalatkerja').trigger("reset");
                  table.draw();
                  $('#saveBtnTambahAlatKerja').html('Tambah');
                  $('#tombolTambahAlatKerja').trigger('click');   
                  $('.table-responsive').before(`<div class="alert alert-warning alert-dismissible fade show" role="alert"> Sukses Update !!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
                  $('.alert').delay(2000).fadeOut(function() {
                     $(this).remove(); 
                  });
                }
        },
        
      });

    })



    //batal
   $('#saveBtnBatalAlatKerja').click( function(e){
      $(this).html('Sending..');
      e.preventDefault();
      $('#formtambahalatkerja').trigger("reset");
         $('#saveBtnBatalAlatKerja').html('Batal');
         $('#tombolTambahAlatKerja').trigger('click');
         $(document).find('span.error-text').text('');
 
    })



    //delete
    $('body').on('click', '.deleteAlatKerja', function () {
     
      var alatkerja_id = $(this).data("id");
      var r = confirm("Are You sure want to delete !");
       if (r == true) {
         $.ajax({
          type: "DELETE",
          url: "/dashboard/master/daftar-alatkerja"+'/'+alatkerja_id,
          success: function (data) {
            $('.table-responsive').before(`<div class="alert alert-danger alert-dismissible fade show" role="alert"> Sukses hapus alatkerja!!<button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>`);
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