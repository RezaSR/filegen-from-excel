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
</interface>
`
)
