	jsoniter "github.com/json-iterator/go"
		{
			name: "ambiguous deleted",
			gitdiff: `diff --git a/b b/b b/b b/b
deleted file mode 100644
index 92e798b..0000000
--- a/b b/b` + "\t" + `
+++ /dev/null
@@ -1 +0,0 @@
-b b/b
`,
			oldFilename: "b b/b",
			filename:    "b b/b",
			addition:    0,
			deletion:    1,
		},
		{
			name: "ambiguous addition",
			gitdiff: `diff --git a/b b/b b/b b/b
new file mode 100644
index 0000000..92e798b
--- /dev/null
+++ b/b b/b` + "\t" + `
@@ -0,0 +1 @@
+b b/b
`,
			oldFilename: "b b/b",
			filename:    "b b/b",
			addition:    1,
			deletion:    0,
		},
		{
			name: "rename",
			gitdiff: `diff --git a/b b/b b/b b/b b/b b/b
similarity index 100%
rename from b b/b b/b b/b b/b
rename to b
`,
			oldFilename: "b b/b b/b b/b b/b",
			filename:    "b",
		},
		{
			name: "ambiguous 1",
			gitdiff: `diff --git a/b b/b b/b b/b b/b b/b
similarity index 100%
rename from b b/b b/b b/b b/b
rename to b
`,
			oldFilename: "b b/b b/b b/b b/b",
			filename:    "b",
		},
		{
			name: "ambiguous 2",
			gitdiff: `diff --git a/b b/b b/b b/b b/b b/b
similarity index 100%
rename from b b/b b/b b/b
rename to b b/b
`,
			oldFilename: "b b/b b/b b/b",
			filename:    "b b/b",
		},
				t.Errorf("ParsePatch(%q) error = %v, wantErr %v", testcase.name, err, testcase.wantErr)
			json := jsoniter.ConfigCompatibleWithStandardLibrary
			gotMarshaled, _ := json.MarshalIndent(got, "", "  ")
				t.Errorf("ParsePath(%q) did not receive 1 file:\n%s", testcase.name, string(gotMarshaled))
				t.Errorf("ParsePath(%q) does not have correct totalAddition %d, wanted %d", testcase.name, got.TotalAddition, testcase.addition)
				t.Errorf("ParsePath(%q) did not have correct totalDeletion %d, wanted %d", testcase.name, got.TotalDeletion, testcase.deletion)
				t.Errorf("ParsePath(%q) does not have correct file addition %d, wanted %d", testcase.name, file.Addition, testcase.addition)
				t.Errorf("ParsePath(%q) did not have correct file deletion %d, wanted %d", testcase.name, file.Deletion, testcase.deletion)
				t.Errorf("ParsePath(%q) did not have correct OldName %q, wanted %q", testcase.name, file.OldName, testcase.oldFilename)
				t.Errorf("ParsePath(%q) did not have correct Name %q, wanted %q", testcase.name, file.Name, testcase.filename)