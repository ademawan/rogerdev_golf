@extends('dashboard.layouts.main')

@section('container')
<div class="row">
<div class="d-flex justify-content-center flex-wrap flex-md-nowrap align-items-center border-bottom">
    <h1 class="h2 d-blockr">DAFTAR MESIN</h1>




    {{-- <div class="btn-toolbar mb-2 mb-md-0">
      <div class="btn-group me-2">
        <button type="button" class="btn btn-sm btn-outline-secondary">Share</button>
        <button type="button" class="btn btn-sm btn-outline-secondary">Export</button>
      </div>
      <button type="button" class="btn btn-sm btn-outline-secondary dropdown-toggle">
        <span data-feather="calendar"></span>
        This week
      </button>
    </div> --}}
  </div>

  {{-- <div class="row">
    {{-- <div class="col">

      <button type="submit">Excel</button>
      <button type="submit">PDF</button>
      <button type="submit">Print</button>
    </div>
    <div class="col">

      <button class="float-end" type="submit">Search</button><input class="float-end" type="text " placeholder="Search....">
     
    </div>
  </div> --}}
  <div class="row my-3">
    <div class="col-md-12">
      <button type="button" id="tombolTambahMesin" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " data-bs-toggle="collapse" data-bs-target="#tambahmesin" aria-expanded="false" aria-controls="collapseExample">
        Tambah Mesin
      </button>
      <div hidden>
        <button type="button" id="tombolRefreshTable" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " >
          Refresh Table
        </button>

      </div>
    </div>
      
    <div class="col-md-12 collapse" id="tambahmesin">
      <div class="col-md">
        <form class="form-horizontal form-disable-enter" name="formtambahmesin" id="formtambahmesin">
          @csrf
          <input type="hidden" id="id_mesin" name="id_mesin" class="form-control" value="">
          <div class="row justify-content-around">
       
            <div class="col-lg-5 shadow-sm">
            
              <div data-spy="scroll" id="scroll-ku" data-bs-target="#scroll-ku" data-bs-offset="0" style="height: 300px;overflow: auto;">
                <div class="row py-2" style="margin:0 !important;">

                  <div class="col-md-12" >
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Jenis Mesin</label>
                      <div class="col-sm-9">
                        <input type="text" id="jenis_mesin" name="jenis_mesin" class="form-control" autofocus>
                        <span class="text-danger error-text jenis_mesin_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Brand</label>
                      <div class="col-sm-9">
                        <input type="text" id="brand" name="brand" class="form-control">
                        <span class="text-danger error-text brand_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Negara Asal</label>
                      <div class="col-sm-9">
                        <input type="text" id="negara_asal" name="negara_asal" class="form-control">
                         <span class="text-danger error-text negara_asal_error"></span>
                      </div>
                    </div>
                  </div>
                  
                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">QTY</label>
                      <div class="col-sm-9">
                        <input type="text" id="qty" name="qty" class="form-control">
                         <span class="text-danger error-text qty_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Tahun</label>
                      <div class="col-sm-9">
                        <input type="date" id="tahun" name="tahun" class="form-control">
                        <span class="text-danger error-text tahun_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">No. Registrasi</label>
                      <div class="col-sm-9">
                        <input type="text" id="no_registrasi" name="no_registrasi" class="form-control">
                          <span class="text-danger error-text no_registrasi_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Daya Kilowatt</label>
                      <div class="col-sm-9">
                        <input type="text" id="daya_kilowatt" name="daya_kilowatt" class="form-control">
                          <span class="text-danger error-text daya_kilowatt_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Daya Voltase</label>
                      <div class="col-sm-9">
                        <input type="text" id="daya_voltase" name="daya_voltase" class="form-control">
                          <span class="text-danger error-text daya_voltase_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Kondisi</label>
                      <div class="col-sm-9">
                        <select class="form-select" id="kondisi" name="kondisi" aria-label="Default select example">
                          <option value="1" selected>Bagus</option>
                          <option value="0" >Rusak</option>
                        </select>
                      </div>
                    </div>      
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Lokasi Mesin</label>
                      <div class="col-sm-9">
                        <input type="text" id="lokasi_mesin" name="lokasi_mesin" class="form-control">
                          <span class="text-danger error-text lokasi_mesin_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Keterangan</label>
                      <div class="col-sm-9">
                        <textarea class="form-control" id="keterangan" name="keterangan" id="floatingTextarea2" style="height: 70px" ></textarea>
                        <span class="text-danger error-text keterangan_error"></span>
                      </div>
                    </div>
                  </div>
                  
                </div>
                <div class="d-grid gap-2 d-md-flex justify-content-md-center telusuri">
                    <button id="saveBtnUpdateMesin" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                    <button id="saveBtnTambahMesin" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                    <button id="saveBtnBatalMesin" class="btn btn-danger updateBtnBatalMesin" type="submit">Batal</button>
                </div>
       
              </div>
            </div>
            <div class="col-lg-6 shadow-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Hic maxime quidem modi repellat fugit velit explicabo nostrum sequi, blanditiis eum voluptas deserunt culpa nemo ut recusandae at nulla illo officiis dolores, cupiditate deleniti vel libero. Accusantium fugit officiis repellat eos sed, voluptatem facere reiciendis eaque inventore cupiditate repellendus veniam veritatis corrupti dignissimos voluptas. Eveniet obcaecati iste dignissimos autem.
              <div class="row">
                <small>
                <span class="text-danger error-text jenis_mesin_errornotif"></span>
                <span class="text-danger error-text brand_errornotif"></span>
                <span class="text-danger error-text negara_asal_errornotif"></span>
                <span class="text-danger error-text qty_errornotif"></span>
                <span class="text-danger error-text tahun_errornotif"></span>
                <span class="text-danger error-text no_registrasi_errornotif"></span>
                <span class="text-danger error-text daya_kilowatt_errornotif"></span>
                <span class="text-danger error-text daya_voltase_errornotif"></span>
                <span class="text-danger error-text kondisi_errornotif"></span>
                <span class="text-danger error-text lokasi_mesin_errornotif"></span>
                <span class="text-danger error-text keterangan_errornotif"></span>
                </small>
              </div>
            </div>
            
          </div>

          </div>

      </form>

      </div>
    </div>
    
  </div>
  <hr class="dropdown-divider mb-3">
 
    <div class="table-responsive">
      <table id="myTable" class="table table-striped" width="100%">
        <thead>
          <tr>
            <th>Jenis Mesin</th>
            <th>Brand</th>
            <th>Negara Asal</th>  
            <th>QTY</th>  
            <th>Tahun</th>  
            <th>No. Registrasi</th>  
            <th>Daya Kilowatt</th>  
            <th>Daya Voltase</th>  
            <th>Kondisi</th>            
            <th>Lokasi Mesin</th>
            <th>Keterangan</th>
            <th>Action</th>
          </tr>     
        </thead>
        <tbody>
          
        </tbody>
      </table>
    </div>


</div>    
@endsection

@push('script')
<script src="/js/master/mesin.js"></script>
    
@endpush