package proggtk 

const (
	gladeStr = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.22.2 -->
<interface>
  <requires lib="gtk+" version="3.20"/>
  <object class="GtkImage" id="image_HelpOutFileName">
    <property name="visible">True</property>
    <property name="can_focus">False</property>
    <property name="stock">gtk-dialog-question</property>
    <property name="icon_size">3</property>
  </object>
  <object class="GtkImage" id="image_HelpTemplateFile">
    <property name="visible">True</property>
    <property name="can_focus">False</property>
    <property name="stock">gtk-dialog-question</property>
    <property name="icon_size">3</property>
  </object>
  <object class="GtkApplicationWindow" id="AppWindow_Main">
    <property name="can_focus">False</property>
    <property name="title" translatable="yes">FileGen from Excel</property>
    <child type="titlebar">
      <placeholder/>
    </child>
    <child>
      <object class="GtkBox" id="Box_Main">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <property name="margin_bottom">10</property>
        <property name="orientation">vertical</property>
        <property name="spacing">10</property>
        <child>
          <object class="GtkMenuBar">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <child>
              <object class="GtkMenuItem">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="label" translatable="yes">_File</property>
                <property name="use_underline">True</property>
                <child type="submenu">
                  <object class="GtkMenu">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <child>
                      <object class="GtkImageMenuItem">
                        <property name="label">gtk-quit</property>
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="use_underline">True</property>
                        <property name="use_stock">True</property>
                        <property name="always_show_image">True</property>
                        <signal name="activate" handler="Menu_Quit_act" swapped="no"/>
                      </object>
                    </child>
                  </object>
                </child>
              </object>
            </child>
            <child>
              <object class="GtkMenuItem">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="label" translatable="yes">_Help</property>
                <property name="use_underline">True</property>
                <child type="submenu">
                  <object class="GtkMenu">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <child>
                      <object class="GtkImageMenuItem">
                        <property name="label">gtk-about</property>
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="use_underline">True</property>
                        <property name="use_stock">True</property>
                        <property name="always_show_image">True</property>
                        <signal name="activate" handler="Menu_About_act" swapped="no"/>
                      </object>
                    </child>
                  </object>
                </child>
              </object>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkGrid" id="Grid_Main">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="margin_left">10</property>
            <property name="margin_right">10</property>
            <property name="row_spacing">10</property>
            <property name="column_spacing">10</property>
            <property name="row_homogeneous">True</property>
            <child>
              <object class="GtkLabel" id="Lbl_DataFile">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Data File</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkFileChooserButton" id="FileBtn_DataFile">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="hexpand">True</property>
                <property name="create_folders">False</property>
                <property name="show_hidden">True</property>
                <property name="title" translatable="yes"/>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="Lbl_TemplateFile">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Template File</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkFileChooserButton" id="FileBtn_TemplateFile">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="hexpand">True</property>
                <property name="create_folders">False</property>
                <property name="show_hidden">True</property>
                <property name="title" translatable="yes"/>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="Btn_HelpTemplateFile">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <property name="halign">center</property>
                <property name="image">image_HelpTemplateFile</property>
                <property name="always_show_image">True</property>
                <signal name="clicked" handler="Btn_HelpTemplateFile_clk" swapped="no"/>
              </object>
              <packing>
                <property name="left_attach">2</property>
                <property name="top_attach">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="Lbl_OutDir">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Output Dir</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkFileChooserButton" id="FileBtn_OutDir">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="hexpand">True</property>
                <property name="action">select-folder</property>
                <property name="show_hidden">True</property>
                <property name="title" translatable="yes"/>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel" id="Lbl_OutFileName">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="halign">end</property>
                <property name="label" translatable="yes">Output File Name</property>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkEntry" id="Entry_OutFileName">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="hexpand">True</property>
              </object>
              <packing>
                <property name="left_attach">1</property>
                <property name="top_attach">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="Btn_HelpOutFileName">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <property name="halign">center</property>
                <property name="image">image_HelpOutFileName</property>
                <property name="always_show_image">True</property>
                <signal name="clicked" handler="Btn_HelpOutFileName_clk" swapped="no"/>
              </object>
              <packing>
                <property name="left_attach">2</property>
                <property name="top_attach">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkButtonBox">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="layout_style">center</property>
                <child>
                  <object class="GtkButton" id="Btn_Generate">
                    <property name="label" translatable="yes">Generate</property>
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <signal name="clicked" handler="Btn_Generate_clk" swapped="no"/>
                  </object>
                  <packing>
                    <property name="expand">True</property>
                    <property name="fill">True</property>
                    <property name="position">0</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="left_attach">0</property>
                <property name="top_attach">4</property>
                <property name="width">3</property>
              </packing>
            </child>
            <child>
              <placeholder/>
            </child>
            <child>
              <placeholder/>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
  <object class="GtkAboutDialog" id="AboutDialog_Main">
    <property name="can_focus">False</property>
    <property name="window_position">center-on-parent</property>
    <property name="type_hint">dialog</property>
    <property name="gravity">center</property>
    <property name="transient_for">AppWindow_Main</property>
    <property name="program_name">FileGen from Excel</property>
    <property name="copyright" translatable="yes">rsr_reza@yahoo.com</property>
    <property name="comments" translatable="yes">Generate text files from each row of Excel data based on a template.</property>
    <property name="website">https://github.com/RezaSR/filegen-from-excel</property>
    <property name="website_label" translatable="yes">View on Github</property>
    <property name="authors">Reza Saberi Rad &lt;rsr_reza@yahoo.com&gt;</property>
    <property name="logo_icon_name">emblem-documents</property>
    <property name="license_type">mit-x11</property>
    <child type="titlebar">
      <placeholder/>
    </child>
    <child internal-child="vbox">
      <object class="GtkBox">
        <property name="can_focus">False</property>
        <property name="orientation">vertical</property>
        <property name="spacing">2</property>
        <child internal-child="action_area">
          <object class="GtkButtonBox">
            <property name="can_focus">False</property>
            <property name="layout_style">end</property>
            <child>
              <placeholder/>
            </child>
            <child>
              <placeholder/>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">False</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <placeholder/>
        </child>
      </object>
    </child>
  </object>
  <object class="GtkDialog" id="Dialog_Main">
    <property name="can_focus">False</property>
    <property name="resizable">False</property>
    <property name="window_position">center-on-parent</property>
    <property name="type_hint">dialog</property>
    <property name="deletable">False</property>
    <property name="gravity">center</property>
    <property name="transient_for">AppWindow_Main</property>
    <child type="titlebar">
      <placeholder/>
    </child>
    <child internal-child="vbox">
      <object class="GtkBox">
        <property name="can_focus">False</property>
        <property name="margin_left">10</property>
        <property name="margin_right">10</property>
        <property name="margin_top">10</property>
        <property name="margin_bottom">10</property>
        <property name="orientation">vertical</property>
        <property name="spacing">10</property>
        <child internal-child="action_area">
          <object class="GtkButtonBox">
            <property name="can_focus">False</property>
            <property name="hexpand">True</property>
            <property name="layout_style">center</property>
            <child>
              <object class="GtkButton" id="Btn_DialogMain">
                <property name="label" translatable="yes">OK</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <signal name="clicked" handler="Btn_DialogMain_clk" swapped="no"/>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">False</property>
            <property name="position">2</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel" id="Lbl_DialogMainHeader">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="selectable">True</property>
            <attributes>
              <attribute name="weight" value="bold"/>
            </attributes>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel" id="Lbl_DialogMainText">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="halign">start</property>
            <property name="selectable">True</property>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>
`
)
