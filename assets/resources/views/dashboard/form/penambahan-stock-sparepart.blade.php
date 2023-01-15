@extends('dashboard.layouts.main')

@section('container')
<div class="row">
<div class="d-flex justify-content-center flex-wrap flex-md-nowrap align-items-center border-bottom">
    <h1 class="h2 d-blockr">PENAMBAHAN STOCK SPAREPART</h1>




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
      <button type="button" id="tombolTambahStockSparepart" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " data-bs-toggle="collapse" data-bs-target="#tambahstocksparepart" aria-expanded="false" aria-controls="collapseExample">
        Tambah Stock Sparepart
      </button>
      <div hidden>
        <button type="button" id="tombolRefreshTable" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " >
          Refresh Table
        </button>

      </div>
    </div>
      
    <div class="col-md-12 collapse" id="tambahstocksparepart">
      <div class="col-md">
        <form class="form-horizontal form-disable-enter" name="formtambahstocksparepart" id="formtambahstocksparepart">
          @csrf
          <input type="hidden" id="id_stocksparepart" name="id_stocksparepart" class="form-control" value="">
          <input type="hidden" id="id_user" name="id_user" class="form-control" value="{{ auth()->user()->id }}">
          <div class="row justify-content-around">
       
            <div class="col-lg-5 shadow-sm">
            
              <div data-spy="scroll" id="scroll-ku" data-bs-target="#scroll-ku" data-bs-offset="0" style="height: 300px;overflow: auto;">
                <div class="row py-2" style="margin:0 !important;">

                  <div class="col-md-12" >
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">Nama Sparepart</label>
                      <div class="col-sm-9">
                        <div class="input-group">
                          <select type="text" data-width="100%" id="selectsparepart" name="selectsparepart" class="form-control" autocomplete="off">
                          
                            <option value="">Pilih Sparepart</option>
                            @foreach ($spareparts as $sparepart)
                                
                            <option value="{{ $sparepart->id }}">{{ $sparepart->nama }}</option>
                            @endforeach
                            
                          </select>
                          <span class="text-danger error-text selectsparepart_error"></span>
  
                        </div>
                        </div>
                    </div>
                  </div>
                  
                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">QTY</label>
                      <div class="col-sm-9">
                        <input type="text" id="qty" name="qty" class="form-control" autocomplete="off">
                         <span class="text-danger error-text qty_error"></span>
                      </div>
                    </div>
                  </div>

                  <div class="col-md-12">
                    <div class="mb-3 row">
                      <label for="inputtext" class="col-sm-3 col-form-label">No. Invoice</label>
                      <div class="col-sm-9">
                        <input type="text" id="no_invoice" name="no_invoice" class="form-control">
                          <span class="text-danger error-text no_invoice_error"></span>
                      </div>
                    </div>
                  </div>
                  
                </div>
                <div class="d-grid gap-2 d-md-flex justify-content-md-center telusuri">
                    <button id="saveBtnUpdateStockSparepart" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                    <button id="saveBtnTambahStockSparepart" class="btn btn-primary me-md-2" type="submit">Tambah</button>
                    <button id="saveBtnBatalStockSparepart" class="btn btn-danger updateBtnBatalStockSparepart" type="submit">Batal</button>
                </div>
       
              </div>
            </div>
            <div class="col-lg-6 shadow-sm">Lorem ipsum dolor sit amet consectetur adipisicing elit. Hic maxime quidem modi repellat fugit velit explicabo nostrum sequi, blanditiis eum voluptas deserunt culpa nemo ut recusandae at nulla illo officiis dolores, cupiditate deleniti vel libero. Accusantium fugit officiis repellat eos sed, voluptatem facere reiciendis eaque inventore cupiditate repellendus veniam veritatis corrupti dignissimos voluptas. Eveniet obcaecati iste dignissimos autem.
              <div class="row">
                <small>
                <span class="text-danger error-text qty_errornotif"></span>
                <span class="text-danger error-text no_invoice_errornotif"></span>
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
            <th>Nama Sparepart</th>
            <th>Nama User</th>  
            <th>Qty</th>  
            <th>No Invoice</th>  
            <th>Tanggal</th>  
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
<script src="/js/form/stock_sparepart.js"></script>
<script src="https://cdn.jsdelivr.net/npm/select2@4.1.0-rc.0/dist/js/select2.min.js"></script>
    
@endpush