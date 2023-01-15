@extends('dashboard.layouts.main')

@section('container')
<div class="row">
  <div class="d-flex justify-content-center flex-wrap flex-md-nowrap align-items-center border-bottom">
    <h1 class="h2 d-blockr">DAFTAR ALAT KERJA</h1>
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
  
    <div class="row my-3">
      <div class="col-md-12">
        <button type="button" id="tombolTambahAlatKerja" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " data-bs-toggle="collapse" data-bs-target="#tambahalatkerja" aria-expanded="false" aria-controls="collapseExample">
          Tambah Alat Kerja
        </button>
        <div hidden>
          <button type="button" id="tombolRefreshTable" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " >
            Refresh Table
          </button>
      </div>
        
      <div class="col-md-12 collapse" id="tambahalatkerja">
        <div class="col-md">
          <form class="form-horizontal form-disable-enter" name="formtambahalatkerja" id="formtambahalatkerja">
            @csrf
            <input type="hidden" id="id_alatkerja" name="id_alatkerja" class="form-control" value="">

            <div class="row justify-content-around">
         
              <div class="col-lg-5 shadow-sm">
              
                <div data-spy="scroll" data-bs-target="#scroll-ku" data-bs-offset="0" style="height: 300px;overflow: auto;">
                  <div class="row py-2" style="margin:0 !important;">
                    <div class="col-md-12" >
                      <div class="mb-3 row">
                        <label for="inputtext" class="col-sm-3 col-form-label">Nama Alat</label>
                        <div class="col-sm-9">
                          <input type="text" name="nama" class="form-control" id="nama">
                          <span class="text-danger error-text nama_error"></span>

                        </div>
                      </div>
                    </div>
                    <div class="col-md-12" >
                        <div class="mb-3 row">
                          <label for="inputtext" class="col-sm-3 col-form-label">Ukuran </label>
                          <div class="col-sm-6">
                              <input type="text" name="ukuran" id="ukuran" class="form-control" aria-label="Text ">
                              <span class="text-danger error-text ukuran_error"></span>
                          </div>
                          <div class="col-sm-3"  style="padding-left: 0 !important">
                            <select class="form-select bg-secondary text-light" name="satuan" id="satuan">
                              <option value="1" selected>mm</option>
                              <option value="2">cm</option>
                              <option value="3">dm</option>
                              <option value="4">m</option>
                              <option value="5">dam</option>
                              <option value="6">hm</option>
                              <option value="7">km</option>
                            </select>
                            
                          </div>  
                        </div>
                      </div>
                   
                    <div class="col-md-12">
                      <div class="mb-3 row">
                        <label for="inputtext" class="col-sm-3 col-form-label">Merk</label>
                        <div class="col-sm-9">  
                            <input type="text" name="merk" class="form-control" id="merk">
                             <span class="text-danger error-text merk_error"></span>
                        </div>
                      </div>      
                    </div>

                    <div class="col-md-12" >
                        <div class="mb-3 row">
                          <label for="inputtext" class="col-sm-3 col-form-label">QTY </label>
                          <div class="col-sm-6">
                              <input type="text" id="qty" name="qty" class="form-control" aria-label="Text ">
                              <span class="text-danger error-text qty_error"></span>
                          </div>
                          <div class="col-sm-3"  style="padding-left: 0 !important">
                            <select class="form-select bg-secondary text-light" id="unit" name="unit">
                              <option value="1" selected>PCS</option>
                              <option value="2">SET</option>
                              <option value="3">BKS</option>
                              <option value="4">ROLL</option>
                              <option value="5">MTR</option>
                              <option value="6">LBR</option>
                            </select>
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
                            <span class="text-danger error-text kondisi_error"></span>
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

                      <div class="col-md-12">
                        <div class="mb-3 row">
                          <label for="inputtext" class="col-sm-3 col-form-label">No Registrasi</label>
                          <div class="col-sm-9">  
                              <input type="text" name="no_registrasi" class="form-control" id="no_registrasi">
                               <span class="text-danger error-text no_registrasi_error"></span>
                          </div>
                        </div>      
                      </div>

                  </div>
                      <div class="d-grid gap-2 d-md-flex justify-content-md-center telusuri">
                        <button id="saveBtnUpdateAlatKerja" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                        <button id="saveBtnTambahAlatKerja" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                        <button id="saveBtnBatalAlatKerja" class="btn btn-danger" type="submit">Batal</button>
                    </div>
         
                </div>
              </div>
              <div class="col-lg-6 shadow-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Hic maxime quidem modi repellat fugit velit explicabo nostrum sequi, blanditiis eum voluptas deserunt culpa nemo ut recusandae at nulla illo officiis dolores, cupiditate deleniti vel libero. Accusantium fugit officiis repellat eos sed, voluptatem facere reiciendis eaque inventore cupiditate repellendus veniam veritatis corrupti dignissimos voluptas. Eveniet obcaecati iste dignissimos autem.
                <div class="row">
                  <small>
                  <span class="text-danger error-text nama_errornotif"></span>
                  <span class="text-danger error-text ukuran_errornotif"></span>
                  <span class="text-danger error-text merk_errornotif"></span>
                  <span class="text-danger error-text qty_errornotif"></span>
                  <span class="text-danger error-text kondisi_errornotif"></span>
                  <span class="text-danger error-text keterangan_errornotif"></span>
                  <span class="text-danger error-text no_registrasi_errornotif"></span>
                  </small>
                </div>

              </div>
            </div>
  
            </div>
  
        </form>
  
        </div>
      </div>
  
    
    <div class="table-responsive">
      <table id="myTable" class="table table-striped w-100">
        <thead>
          <tr>
            <th scope="col">Nama Alat</th>
            <th scope="col">Ukuran</th>
            <th scope="col">Merk</th>
            <th scope="col">Qty</th>
            <th scope="col">Kondisi</th>
            <th scope="col">Keterangan</th>
            <th scope="col">No Registrasi</th>
            <th scope="col">Action</th>
          </tr>
        </thead>
        <tbody>
          
        </tbody>
      </table>
    </div>
      

</div>

@endsection

@push('script')
<script type="text/javascript" src="/js/master/alat_kerja.js"></script>
    
@endpush