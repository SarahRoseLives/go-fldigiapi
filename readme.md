# FLDIGI XMLRPC API WRITTEN IN GO

## This is a work in progress application.
We currently support only string and nil type responses and arguments. I'm very lacking when it comes to non-dynamicly typed lanaguages. I hope to fully support the api in the future.


## XML Command Symbol Interpretation

This document provides an interpretation guide for symbols used in XML commands and presents a table detailing various XML commands along with their method names, signatures, and descriptions.

### Working Methods are in Bold 

## Symbol Interpretation

| Symbol | Interpretation |
| ------ | -------------- |
| 6      | bytes          |
| A      | array          |
| b      | boolean        |
| d      | double         |
| i      | integer        |
| n      | nil            |
| s      | string         |
| S      | struct         |

## Table of XML Commands

| Method Name             | Sig (ret:arg) | Description                                                  |
| ----------------------- | ------------- | ------------------------------------------------------------ |
| <b>fldigi.config_dir       | s:n           | Returns the name of the configuration directory              |
| fldigi.list             | A:n           | Returns the list of methods                                   |
| <b>fldigi.name             | s:n           | Returns the program name                                      |
| <b>fldigi.name_version     | s:n           | Returns the program name and version                          |
| fldigi.terminate        | n:i           | Terminates fldigi. `'i'` is bitmask specifying data to save: 0=options; 1=log; 2=macros |
| <b>fldigi.version          | s:n           | Returns the program version as a string                       |
| fldigi.version_struct   | S:n           | Returns the program version as a struct                       |
| <b>io.enable_arq           | n:n           | Switch to ARQ I/O                                            |
| <b>io.enable_kiss          | n:n           | Switch to KISS I/O                                           |
| <b>io.in_use               | s:n           | Returns the IO port in use (ARQ/KISS).                        |
| <b>log.clear               | n:n           | Clears the contents of the log fields                         |
| <b>log.get_az              | s:n           | Returns the AZ field contents                                 |
| <b>log.get_band            | s:n           | Returns the current band name                                 |
| <b>log.get_call            | s:n           | Returns the Call field contents                               |
| <b>log.get_country         | s:n           | Returns the Country field contents                            |
| <b>log.get_exchange        | s:n           | Returns the contest exchange field contents                   |
| <b>log.get_frequency       | s:n           | Returns the Frequency field contents                          |
| <b>log.get_locator         | s:n           | Returns the Locator field contents                            |
| <b>log.get_name            | s:n           | Returns the Name field contents                               |
| <b>log.get_notes           | s:n           | Returns the Notes field contents                              |
| <b>log.get_province        | s:n           | Returns the Province field contents                           |
| <b>log.get_qth             | s:n           | Returns the QTH field contents                                |
| <b>log.get_rst_in          | s:n           | Returns the RST(r) field contents                             |
| <b>log.get_rst_out         | s:n           | Returns the RST(s) field contents                             |
| <b>log.get_serial_number   | s:n           | Returns the serial number field contents                      |
| <b>log.get_serial_number_sent | s:n         | Returns the serial number (sent) field contents               |
| <b>log.get_state           | s:n           | Returns the State field contents                              |
| <b>log.get_time_off        | s:n           | Returns the Time-Off field contents                           |
| <b>log.get_time_on         | s:n           | Returns the Time-On field contents                            |
| <b>log.set_call            | n:s           | Sets the Call field contents                                  |
| <b>log.set_exchange        | n:s           | Sets the contest exchange field contents                      |
| <b>log.set_locator         | n:s           | Sets the Locator field contents                               |
| <b>log.set_name            | n:s           | Sets the Name field contents                                  |
| <b>log.set_qth             | n:s           | Sets the QTH field contents                                   |
| <b>log.set_rst_in          | n:s           | Sets the RST(r) field contents                                |
| <b>log.set_rst_out         | n:s           | Sets the RST(s) field contents                                |
| <b>log.set_serial_number   | n:s           | Sets the serial number field contents                         |
| <b>main.abort              | n:n           | Aborts a transmit or tune                                     |
| main.get_afc            | b:n           | Returns the AFC state                                         |
| <b>main.get_char_rates     | s:n           | Returns table of char rates.                                  |
| main.get_char_timing    | n:6           | Input: value of character. Returns transmit duration for specified character (samples:sample rate). |
| main.get_frequency      | d:n           | Returns the RF carrier frequency                              |
| main.get_lock           | b:n           | Returns the Transmit Lock state                               |
| main.get_max_macro_id   | i:n           | Returns the maximum macro ID number                           |
| main.get_reverse        | b:n           | Returns the Reverse Sideband state                            |
| main.get_rsid           | b:n           | Returns the RSID state                                        |
| main.get_txid           | b:n           | Returns the TxRSID state                                      |
| main.get_squelch        | b:n           | Returns the squelch state                                     |
| main.get_squelch_level  | d:n           | Returns the squelch level                                     |
|<b> main.get_status1        | s:n           | Returns the contents of the first status field (typically s/n) |
| <b>main.get_status2        | s:n           | Returns the contents of the second status field               |
| <b>main.get_trx_state      | s:n           | Returns T/R state                                             |
| <b>main.get_trx_status     | s:n           | Returns transmit/tune/receive status                          |
| main.get_tx_timing      | n:6           | Returns transmit duration for test string (samples:sample rate:secs). |
| <b>main.get_wf_sideband    | s:n           | Returns the current waterfall sideband                        |
| main.inc_frequency      | d:d           | Increments the RF carrier frequency. Returns the new value    |
| main.inc_squelch_level  | d:d           | Increments the squelch level. Returns the new level           |
| main.run_macro          | n:i           | Runs a macro                                                 |
| <b>main.rx                 | n:n           | Receives                                                     |
| <b>main.rx_only            | n:n           | Disables Tx.                                                 |
| <b>main.rx_tx              | n:n           | Sets normal Rx/Tx switching.                                 |
| main.set_afc            | b:b           | Sets the AFC state. Returns the old state                    |
| main.set_frequency      | d:d           | Sets the RF carrier frequency. Returns the old value          |
| main.set_lock           | b:b           | Sets the Transmit Lock state. Returns the old state           |
| main.set_reverse        | b:b           | Sets the Reverse Sideband state. Returns the old state        |
| main.set_rsid           | b:b           | Sets the RSID state. Returns the old state                    |
| main.set_txid           | b:b           | Sets the TxRSID state. Returns the old state                   |
| main.set_squelch        | b:b           | Sets the squelch state. Returns the old state                 |
| main.set_squelch_level  | d:d           | Sets the squelch level. Returns the old level                 |
| main.set_wf_sideband    | n:s           | Sets the waterfall sideband to USB or LSB                     |
| main.toggle_afc        | b:n           | Toggles the AFC state. Returns the new state                  |
| main.toggle_lock       | b:n           | Toggles the Transmit Lock state. Returns the new state        |
| main.toggle_reverse     | b:n           | Toggles the Reverse Sideband state. Returns the new state     |
| main.toggle_rsid       | b:n           | Toggles the RSID state. Returns the new state                 |
| main.toggle_txid       | b:n           | Toggles the TxRSID state. Returns the new state               |
| main.toggle_squelch    | b:n           | Toggles the squelch state. Returns the new state              |
| <b>main.tune               | n:n           | Tunes                                                        |
| <b>main.tx                 | n:n           | Transmits                                                    |
| modem.get_afc_search_range | i:n        | Returns the modem AFC search range                            |
| modem.get_bandwidth     | i:n           | Returns the modem bandwidth                                   |
| modem.get_carrier       | i:n           | Returns the modem carrier frequency                           |
| modem.get_id            | i:n           | Returns the ID of the current modem                           |
| modem.get_max_id        | i:n           | Returns the maximum modem ID number                           |
| <b>modem.get_name          | s:n           | Returns the name of the current modem                         |
| modem.get_names         | A:n           | Returns all modem names                                       |
| modem.get_quality       | d:n           | Returns the modem signal quality in the range [0:100]         |
| modem.inc_afc_search_range | n:i        | Increments the modem AFC search range. Returns the new value   |
| modem.inc_bandwidth     | n:i           | Increments the modem bandwidth. Returns the new value          |
| modem.inc_carrier       | i:i           | Increments the modem carrier frequency. Returns the new carrier |
| modem.olivia.get_bandwidth | i:n        | Returns the Olivia bandwidth                                  |
| modem.olivia.get_tones  | i:n           | Returns the Olivia tones                                      |
| modem.olivia.set_bandwidth | n:i         | Sets the Olivia bandwidth                                    |
| modem.olivia.set_tones  | n:i           | Sets the Olivia tones                                         |
| <b>modem.search_down       | n:n           | Searches downward in frequency                                |
| <b>modem.search_up         | n:n           | Searches upward in frequency                                  |
| modem.set_afc_search_range | n:i        | Sets the modem AFC search range. Returns the old value         |
| modem.set_bandwidth     | n:i           | Sets the modem bandwidth. Returns the old value                |
| modem.set_by_id         | i:i           | Sets the current modem. Returns old ID                        |
| <b>modem.set_by_name       | s:s           | Sets the current modem. Returns old name                      |
| modem.set_carrier       | i:i           | Sets modem carrier. Returns old carrier                       |
| navtex.get_message      | s:i           | Returns next Navtex/SitorB message with a max delay in seconds.. Empty string if timeout. |
| <b>navtex.send_message     | s:s           | Send a Navtex/SitorB message. Returns an empty string if OK otherwise an error message. |
| <b>rig.get_bandwidth       | s:n           | Returns the name of the current transceiver bandwidth         |
| rig.get_bandwidths      | A:n           | Returns the list of available rig bandwidths                  |
| <b>rig.get_mode            | s:n           | Returns the name of the current transceiver mode              |
| rig.get_modes           | A:n           | Returns the list of available rig modes                       |
| <b>rig.get_name            | s:n           | Returns the rig name previously set via rig.set_name          |
| <b>rig.get_notch           | s:n           | Reports a notch filter frequency based on WF action            |
| <b>rig.release_control     | n:n           | Switches rig control to previous setting                      |
| <b>rig.set_bandwidth       | n:s           | Selects a bandwidth previously added by rig.set_bandwidths    |
| rig.set_bandwidths      | n:A           | Sets the list of available rig bandwidths                     |
| rig.set_frequency       | d:d           | Sets the RF carrier frequency. Returns the old value          |
| <b>rig.set_mode            | n:s           | Selects a mode previously added by rig.set_modes              |
| rig.set_modes           | n:A           | Sets the list of available rig modes                          |
| <b>rig.set_name            | n:s           | Sets the rig name for xmlrpc rig                              |
| rig.set_pwrmeter        | n:i           | Sets the power meter returns null.                             |
| rig.set_smeter          | n:i           | Sets the smeter returns null.                                  |
| <b>rig.take_control        | n:n           | Switches rig control to XML-RPC                                |
| rx.get_data             | 6:n           | Returns all RX data received since last query.                |
| rxtx.get_data           | 6:n           | Returns all RXTX combined data since last query.              |
| spot.get_auto           | b:n           | Returns the autospotter state                                 |
| spot.pskrep.get_count   | i:n           | Returns the number of callsigns spotted in the current session |
| spot.set_auto           | n:b           | Sets the autospotter state. Returns the old state             |
| spot.toggle_auto        | n:b           | Toggles the autospotter state. Returns the new state           |
| <b>text.add_tx             | n:s           | Adds a string to the TX text widget                           |
| text.add_tx_bytes       | n:6           | Adds a byte string to the TX text widget                      |
| <b>text.clear_rx           | n:n           | Clears the RX text widget                                     |
| <b>text.clear_tx           | n:n           | Clears the TX text widget                                     |
| text.get_rx             | 6:i           | Returns a range of characters (start, length) from the RX text widget |
| text.get_rx_length      | i:n           | Returns the number of characters in the RX widget             |
| tx.get_data             | 6:n           | Returns all TX data transmitted since last query.             |
| <b>wefax.end_reception     | s:n           | End Wefax image reception                                    |
| wefax.get_received_file | s:i           | Waits for next received fax file, returns its name with a delay. Empty string if timeout. |
| wefax.send_file         | s:i           | Send file. returns an empty string if OK otherwise an error message. |
| wefax.set_adif_log      | s:b           | Set/reset logging to received/transmit images to ADIF log file |
| wefax.set_max_lines     | s:i           | Set maximum lines for fax image reception                     |
| <b>wefax.set_tx_abort_flag | s:n           | Cancels Wefax image transmission                             |
| <b>wefax.skip_apt          | s:n           | Skip APT during Wefax reception                               |
| <b>wefax.skip_phasing      | s:n           | Skip phasing during Wefax reception                            |
| <b>wefax.state_string      | s:n           | Returns Wefax engine state (tx and rx) for information.       |

