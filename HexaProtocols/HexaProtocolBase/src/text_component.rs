use serde_json::json; // Usaremos serde_json para facilitar la creación de JSON

#[derive(Debug, Default)]
pub struct TextComponent {
    pub text: Option<String>,
    pub translate: Option<String>,
    pub with: Option<Vec<TextComponent>>,
    pub extra: Option<Vec<TextComponent>>,
    pub color: Option<String>,
    pub bold: Option<bool>,
    pub italic: Option<bool>,
    pub underlined: Option<bool>,
    pub strikethrough: Option<bool>,
    pub obfuscated: Option<bool>,
    pub font: Option<String>,
    pub insertion: Option<String>,
}

impl TextComponent {
    pub fn new() -> Self {
        TextComponent::default()
    }

    pub fn set_text(&mut self, text: &str) -> &mut Self {
        self.text = Some(text.to_string().replace("&", "§"));
        self
    }

    pub fn set_translate(&mut self, key: &str) -> &mut Self {
        self.translate = Some(key.to_string());
        self
    }

    pub fn add_with(&mut self, component: TextComponent) -> &mut Self {
        if self.with.is_none() {
            self.with = Some(Vec::new());
        }
        self.with.as_mut().unwrap().push(component);
        self
    }

    pub fn add_extra(&mut self, component: TextComponent) -> &mut Self {
        if self.extra.is_none() {
            self.extra = Some(Vec::new());
        }
        self.extra.as_mut().unwrap().push(component);
        self
    }

    pub fn set_color(&mut self, color: &str) -> &mut Self {
        self.color = Some(color.to_string());
        self
    }

    pub fn set_bold(&mut self, bold: bool) -> &mut Self {
        self.bold = Some(bold);
        self
    }

    pub fn set_italic(&mut self, italic: bool) -> &mut Self {
        self.italic = Some(italic);
        self
    }

    pub fn set_underlined(&mut self, underlined: bool) -> &mut Self {
        self.underlined = Some(underlined);
        self
    }

    pub fn set_strikethrough(&mut self, strikethrough: bool) -> &mut Self {
        self.strikethrough = Some(strikethrough);
        self
    }

    pub fn set_obfuscated(&mut self, obfuscated: bool) -> &mut Self {
        self.obfuscated = Some(obfuscated);
        self
    }

    pub fn set_font(&mut self, font: &str) -> &mut Self {
        self.font = Some(font.to_string());
        self
    }

    pub fn set_insertion(&mut self, insertion: &str) -> &mut Self {
        self.insertion = Some(insertion.to_string());
        self
    }

    pub fn to_json(&self) -> serde_json::Value {
        let mut component = json!({});

        if let Some(ref text) = self.text {
            component["text"] = json!(text);
        }

        if let Some(ref translate) = self.translate {
            component["translate"] = json!(translate);
        }

        if let Some(ref with) = self.with {
            component["with"] = json!(with.iter().map(|c| c.to_json()).collect::<Vec<_>>());
        }

        if let Some(ref extra) = self.extra {
            component["extra"] = json!(extra.iter().map(|c| c.to_json()).collect::<Vec<_>>());
        }

        if let Some(ref color) = self.color {
            component["color"] = json!(color);
        }

        if let Some(bold) = self.bold {
            component["bold"] = json!(bold);
        }

        if let Some(italic) = self.italic {
            component["italic"] = json!(italic);
        }

        if let Some(underlined) = self.underlined {
            component["underlined"] = json!(underlined);
        }

        if let Some(strikethrough) = self.strikethrough {
            component["strikethrough"] = json!(strikethrough);
        }

        if let Some(obfuscated) = self.obfuscated {
            component["obfuscated"] = json!(obfuscated);
        }

        if let Some(ref font) = self.font {
            component["font"] = json!(font);
        }

        if let Some(ref insertion) = self.insertion {
            component["insertion"] = json!(insertion);
        }

        component
    }
}