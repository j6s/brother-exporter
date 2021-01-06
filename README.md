# Brother Exporter

Simple prometheus exporter that exposes printer statistics of my printer.
This exporter is built for a **Brother MFC-L2710DN** but may work with
other printers that expose statistics as a CSV file under `/etc/mnt_info.csv`.

## Usage

```
$ brother-exporter --help
Usage of brother-exporter:
  -listen string
    	Host and port to listen on (default ":9055")
```

## Example prometheus configuration

```

```

## Example of exposed statistics

```
# HELP brother_a4_letter 
# TYPE brother_a4_letter gauge
brother_a4_letter{host="192.168.99.15"} 1022
# HELP brother_a5 
# TYPE brother_a5 gauge
brother_a5{host="192.168.99.15"} 0
# HELP brother_adf_scan 
# TYPE brother_adf_scan gauge
brother_adf_scan{host="192.168.99.15"} 547
# HELP brother_average_coverage 
# TYPE brother_average_coverage gauge
brother_average_coverage{host="192.168.99.15"} 5.2
# HELP brother_b5_executive 
# TYPE brother_b5_executive gauge
brother_b5_executive{host="192.168.99.15"} 0
# HELP brother_copy 
# TYPE brother_copy gauge
brother_copy{host="192.168.99.15"} 14
# HELP brother_copy_2_sided_print 
# TYPE brother_copy_2_sided_print gauge
brother_copy_2_sided_print{host="192.168.99.15"} 0
# HELP brother_envelopes 
# TYPE brother_envelopes gauge
brother_envelopes{host="192.168.99.15"} 0
# HELP brother_envelopes_env_thick_env_thin 
# TYPE brother_envelopes_env_thick_env_thin gauge
brother_envelopes_env_thick_env_thin{host="192.168.99.15"} 0
# HELP brother_error_count_1 
# TYPE brother_error_count_1 gauge
brother_error_count_1{host="192.168.99.15"} 1000
# HELP brother_error_count_10 
# TYPE brother_error_count_10 gauge
brother_error_count_10{host="192.168.99.15"} 0
# HELP brother_error_count_2 
# TYPE brother_error_count_2 gauge
brother_error_count_2{host="192.168.99.15"} 708
# HELP brother_error_count_3 
# TYPE brother_error_count_3 gauge
brother_error_count_3{host="192.168.99.15"} 620
# HELP brother_error_count_4 
# TYPE brother_error_count_4 gauge
brother_error_count_4{host="192.168.99.15"} 380
# HELP brother_error_count_5 
# TYPE brother_error_count_5 gauge
brother_error_count_5{host="192.168.99.15"} 350
# HELP brother_error_count_6 
# TYPE brother_error_count_6 gauge
brother_error_count_6{host="192.168.99.15"} 0
# HELP brother_error_count_7 
# TYPE brother_error_count_7 gauge
brother_error_count_7{host="192.168.99.15"} 0
# HELP brother_error_count_8 
# TYPE brother_error_count_8 gauge
brother_error_count_8{host="192.168.99.15"} 0
# HELP brother_error_count_9 
# TYPE brother_error_count_9 gauge
brother_error_count_9{host="192.168.99.15"} 0
# HELP brother_fax 
# TYPE brother_fax gauge
brother_fax{host="192.168.99.15"} 0
# HELP brother_fax_2_sided_print 
# TYPE brother_fax_2_sided_print gauge
brother_fax_2_sided_print{host="192.168.99.15"} 0
# HELP brother_flatbed_scan 
# TYPE brother_flatbed_scan gauge
brother_flatbed_scan{host="192.168.99.15"} 76
# HELP brother_hagaki 
# TYPE brother_hagaki gauge
brother_hagaki{host="192.168.99.15"} 0
# HELP brother_jam_2_sided 
# TYPE brother_jam_2_sided gauge
brother_jam_2_sided{host="192.168.99.15"} 0
# HELP brother_jam_inside 
# TYPE brother_jam_inside gauge
brother_jam_inside{host="192.168.99.15"} 0
# HELP brother_jam_rear 
# TYPE brother_jam_rear gauge
brother_jam_rear{host="192.168.99.15"} 0
# HELP brother_jam_tray_1 
# TYPE brother_jam_tray_1 gauge
brother_jam_tray_1{host="192.168.99.15"} 0
# HELP brother_label 
# TYPE brother_label gauge
brother_label{host="192.168.99.15"} 0
# HELP brother_legal_folio 
# TYPE brother_legal_folio gauge
brother_legal_folio{host="192.168.99.15"} 0
# HELP brother_memory_size 
# TYPE brother_memory_size gauge
brother_memory_size{host="192.168.99.15"} 64
# HELP brother_others 
# TYPE brother_others gauge
brother_others{host="192.168.99.15"} 4
# HELP brother_others_2_sided_print 
# TYPE brother_others_2_sided_print gauge
brother_others_2_sided_print{host="192.168.99.15"} 2
# HELP brother_page_counter 
# TYPE brother_page_counter gauge
brother_page_counter{host="192.168.99.15"} 1022
# HELP brother_percent_of_life_remaining_drum_unit_ 
# TYPE brother_percent_of_life_remaining_drum_unit_ gauge
brother_percent_of_life_remaining_drum_unit_{host="192.168.99.15"} 92
# HELP brother_percent_of_life_remaining_toner_ 
# TYPE brother_percent_of_life_remaining_toner_ gauge
brother_percent_of_life_remaining_toner_{host="192.168.99.15"} 90
# HELP brother_plain_thin_recycled 
# TYPE brother_plain_thin_recycled gauge
brother_plain_thin_recycled{host="192.168.99.15"} 1022
# HELP brother_print 
# TYPE brother_print gauge
brother_print{host="192.168.99.15"} 1004
# HELP brother_print_2_sided_print 
# TYPE brother_print_2_sided_print gauge
brother_print_2_sided_print{host="192.168.99.15"} 322
# HELP brother_replace_count_drum_unit_ 
# TYPE brother_replace_count_drum_unit_ gauge
brother_replace_count_drum_unit_{host="192.168.99.15"} 0
# HELP brother_replace_count_toner_ 
# TYPE brother_replace_count_toner_ gauge
brother_replace_count_toner_{host="192.168.99.15"} 1
# HELP brother_scan_page_count 
# TYPE brother_scan_page_count gauge
brother_scan_page_count{host="192.168.99.15"} 609
# HELP brother_thick_thicker_bond 
# TYPE brother_thick_thicker_bond gauge
brother_thick_thicker_bond{host="192.168.99.15"} 0
# HELP brother_total 
# TYPE brother_total gauge
brother_total{host="192.168.99.15"} 1022
# HELP brother_total_2_sided_print 
# TYPE brother_total_2_sided_print gauge
brother_total_2_sided_print{host="192.168.99.15"} 324
# HELP brother_total_paper_jams 
# TYPE brother_total_paper_jams gauge
brother_total_paper_jams{host="192.168.99.15"} 0
# HELP brother_total_paper_jams_adf_ 
# TYPE brother_total_paper_jams_adf_ gauge
brother_total_paper_jams_adf_{host="192.168.99.15"} 3

```