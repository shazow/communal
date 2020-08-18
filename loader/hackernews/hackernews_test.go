package hackernews

import (
	"context"
	"net/http"
	"testing"

	"github.com/shazow/communal/internal/httphelper"
)

func TestLinksFromComments(t *testing.T) {
	loader := Loader{
		Client: http.Client{
			Transport: httphelper.FixedRoundTrip(sampleNewsOutput),
		},
	}

	ctx := context.Background()
	res, err := loader.Search(ctx, "foo")
	if err != nil {
		t.Fatal(err)
	}

	loader.Client.Transport = httphelper.FixedRoundTrip(sampleCommentsOutput)

	links, err := loader.linksFromComments(ctx, res)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(links), 80; got != want {
		t.Errorf("got: %d; want: %d", got, want)
	}
}

const sampleNewsOutput = `
{
  "hits": [
    {
      "created_at": "2015-01-02T21:05:46.000Z",
      "title": "Why aren’t we using SSH for everything?",
      "url": "https://medium.com/@shazow/ssh-how-does-it-even-9e43586e4ffc",
      "author": "dkua",
      "points": 270,
      "story_text": "",
      "comment_text": null,
      "num_comments": 103,
      "story_id": null,
      "story_title": null,
      "story_url": null,
      "parent_id": null,
      "created_at_i": 1420232746,
      "relevancy_score": 5766,
      "_tags": [
        "story",
        "author_dkua",
        "story_8828543"
      ],
      "objectID": "8828543",
      "_highlightResult": {
        "title": {
          "value": "Why aren’t we using SSH for everything?",
          "matchLevel": "none",
          "matchedWords": []
        },
        "url": {
          "value": "<em>https</em>:<em>//medium.com</em>/@<em>shazow</em><em>/ssh-how-does-it-even-9e43586e4ffc</em>",
          "matchLevel": "full",
          "fullyHighlighted": true,
          "matchedWords": [
            "https",
            "medium",
            ".",
            "com",
            "shazow",
            "ssh",
            "how",
            "does",
            "it",
            "even",
            "9e43586e4ffc"
          ]
        },
        "author": {
          "value": "dkua",
          "matchLevel": "none",
          "matchedWords": []
        },
        "story_text": {
          "value": "",
          "matchLevel": "none",
          "matchedWords": []
        }
      }
    },
    {
      "created_at": "2016-04-17T22:13:06.000Z",
      "title": "Why aren’t we using SSH for everything? (2015)",
      "url": "https://medium.com/@shazow/ssh-how-does-it-even-9e43586e4ffc#.r9qzj6hhl",
      "author": "fforflo",
      "points": 173,
      "story_text": null,
      "comment_text": null,
      "num_comments": 84,
      "story_id": null,
      "story_title": null,
      "story_url": null,
      "parent_id": null,
      "created_at_i": 1460931186,
      "relevancy_score": 6666,
      "_tags": [
        "story",
        "author_fforflo",
        "story_11516582"
      ],
      "objectID": "11516582",
      "_highlightResult": {
        "title": {
          "value": "Why aren’t we using SSH for everything? (2015)",
          "matchLevel": "none",
          "matchedWords": []
        },
        "url": {
          "value": "<em>https</em>:<em>//medium.com</em>/@<em>shazow</em><em>/ssh-how-does-it-even-9e43586e4ffc</em>#.r9qzj6hhl",
          "matchLevel": "full",
          "fullyHighlighted": false,
          "matchedWords": [
            "https",
            "medium",
            ".",
            "com",
            "shazow",
            "ssh",
            "how",
            "does",
            "it",
            "even",
            "9e43586e4ffc"
          ]
        },
        "author": {
          "value": "fforflo",
          "matchLevel": "none",
          "matchedWords": []
        }
      }
    },
    {
      "created_at": "2017-02-04T12:47:25.000Z",
      "title": "Why aren’t we using SSH for everything? (2015)",
      "url": "https://medium.com/@shazow/ssh-how-does-it-even-9e43586e4ffc#.neyhpij1r",
      "author": "turrini",
      "points": 58,
      "story_text": null,
      "comment_text": null,
      "num_comments": 9,
      "story_id": null,
      "story_title": null,
      "story_url": null,
      "parent_id": null,
      "created_at_i": 1486212445,
      "relevancy_score": 7231,
      "_tags": [
        "story",
        "author_turrini",
        "story_13566801"
      ],
      "objectID": "13566801",
      "_highlightResult": {
        "title": {
          "value": "Why aren’t we using SSH for everything? (2015)",
          "matchLevel": "none",
          "matchedWords": []
        },
        "url": {
          "value": "<em>https</em>:<em>//medium.com</em>/@<em>shazow</em><em>/ssh-how-does-it-even-9e43586e4ffc</em>#.neyhpij1r",
          "matchLevel": "full",
          "fullyHighlighted": false,
          "matchedWords": [
            "https",
            "medium",
            ".",
            "com",
            "shazow",
            "ssh",
            "how",
            "does",
            "it",
            "even",
            "9e43586e4ffc"
          ]
        },
        "author": {
          "value": "turrini",
          "matchLevel": "none",
          "matchedWords": []
        }
      }
    },
    {
      "created_at": "2015-01-03T13:28:02.000Z",
      "title": "Why aren’t we using SSH for everything?",
      "url": "https://medium.com/@shazow/ssh-how-does-it-even-9e43586e4ffc?hn=1",
      "author": "mparramon",
      "points": 3,
      "story_text": "",
      "comment_text": null,
      "num_comments": 0,
      "story_id": null,
      "story_title": null,
      "story_url": null,
      "parent_id": null,
      "created_at_i": 1420291682,
      "relevancy_score": 5764,
      "_tags": [
        "story",
        "author_mparramon",
        "story_8830693"
      ],
      "objectID": "8830693",
      "_highlightResult": {
        "title": {
          "value": "Why aren’t we using SSH for everything?",
          "matchLevel": "none",
          "matchedWords": []
        },
        "url": {
          "value": "<em>https</em>:<em>//medium.com</em>/@<em>shazow</em><em>/ssh-how-does-it-even-9e43586e4ffc</em>?hn=1",
          "matchLevel": "full",
          "fullyHighlighted": true,
          "matchedWords": [
            "https",
            "medium",
            ".",
            "com",
            "shazow",
            "ssh",
            "how",
            "does",
            "it",
            "even",
            "9e43586e4ffc"
          ]
        },
        "author": {
          "value": "mparramon",
          "matchLevel": "none",
          "matchedWords": []
        },
        "story_text": {
          "value": "",
          "matchLevel": "none",
          "matchedWords": []
        }
      }
    }
  ],
  "nbHits": 4,
  "page": 0,
  "nbPages": 1,
  "hitsPerPage": 20,
  "exhaustiveNbHits": true,
  "query": "https://medium.com/@shazow/ssh-how-does-it-even-9e43586e4ffc",
  "params": "advancedSyntax=true&analytics=true&analyticsTags=backend&query=https%3A%2F%2Fmedium.com%2F%40shazow%2Fssh-how-does-it-even-9e43586e4ffc&restrictSearchableAttributes=url",
  "processingTimeMS": 9
}
`

const sampleCommentsOutput = `
{
  "id": 8828543,
  "created_at": "2015-01-02T21:05:46.000Z",
  "created_at_i": 1420232746,
  "type": "story",
  "author": "dkua",
  "title": "Why aren’t we using SSH for everything?",
  "url": "https://medium.com/@shazow/ssh-how-does-it-even-9e43586e4ffc",
  "text": "",
  "points": 270,
  "parent_id": null,
  "story_id": null,
  "children": [
    {
      "id": 8829186,
      "created_at": "2015-01-02T23:36:48.000Z",
      "created_at_i": 1420241808,
      "type": "comment",
      "author": "silbak04",
      "title": null,
      "url": null,
      "text": "<p>Somewhat relevant: I think Mosh solves a few problems that SSH lacks. This article: <a href=\"https://news.ycombinator.com/item?id=8252093\" rel=\"nofollow\">https:&#x2F;&#x2F;news.ycombinator.com&#x2F;item?id=8252093</a> made it on Hacker News a bit ago. The video is wort the watch imo.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [],
      "options": []
    },
    {
      "id": 8830466,
      "created_at": "2015-01-03T10:48:56.000Z",
      "created_at_i": 1420282136,
      "type": "comment",
      "author": "emmanueloga_",
      "title": null,
      "url": null,
      "text": "<p>If you rephrase this question like this, the answer is evident:</p><p>&quot;Why aren&#x27;t we using a protocol designed to add encryption to pseudo-devices emulating a real text terminal device [1] for <i>everything</i>?&quot;</p><p>1: <a href=\"http://en.wikipedia.org/wiki/Pseudo_terminal\" rel=\"nofollow\">http:&#x2F;&#x2F;en.wikipedia.org&#x2F;wiki&#x2F;Pseudo_terminal</a></p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [],
      "options": []
    },
    {
      "id": 8829208,
      "created_at": "2015-01-02T23:43:00.000Z",
      "created_at_i": 1420242180,
      "type": "comment",
      "author": "otterley",
      "title": null,
      "url": null,
      "text": "<p>For one thing, SSH performance for large transfers is abysmal over high-latency links.  SSH uses a small TCP window size because it was optimized for quick response, not bulk transfers.  (After all, it was designed to be a remote shell protocol, not a generic file transfer protocol; scp and sftp were added after the fact.)  And unfortunately you can&#x27;t currently specify a different buffer size even if you wanted to.</p><p>More information can be found here: <a href=\"http://www.psc.edu/index.php/hpn-ssh\" rel=\"nofollow\">http:&#x2F;&#x2F;www.psc.edu&#x2F;index.php&#x2F;hpn-ssh</a></p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829722,
          "created_at": "2015-01-03T02:47:49.000Z",
          "created_at_i": 1420253269,
          "type": "comment",
          "author": "bodyfour",
          "title": null,
          "url": null,
          "text": "<p>&gt; scp and sftp were added after the fact</p><p>scp was there from the beginning, or thereabouts.  It certainly was present the first time I played with it which was within a few months of the first release.</p><p>ssh was specifically designed as a replacement for BSD remote tools (rsh&#x2F;rlogin&#x2F;rcp) even falling back to those protocols if ssh wasn&#x27;t available (although the FallBackToRsh option was removed from OpenSSH a decade or so ago I think).  That&#x27;s also why &quot;slogin&quot; is installed as a symlink to &quot;ssh&quot;, to keep it similar to rsh&#x2F;rlogin.  It only made sense for there to be a &quot;scp&quot; to work like the existing &quot;rcp&quot;.</p><p>You are right that sftp only came later.</p>",
          "points": null,
          "parent_id": 8829208,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8829500,
          "created_at": "2015-01-03T01:06:49.000Z",
          "created_at_i": 1420247209,
          "type": "comment",
          "author": "jzila",
          "title": null,
          "url": null,
          "text": "<p>Are there any plans to patch this into OpenSSH? It looks like they have a great solution for the problem you highlighted.</p>",
          "points": null,
          "parent_id": 8829208,
          "story_id": 8828543,
          "children": [
            {
              "id": 8829731,
              "created_at": "2015-01-03T02:49:55.000Z",
              "created_at_i": 1420253395,
              "type": "comment",
              "author": "dmm",
              "title": null,
              "url": null,
              "text": "<p>Those patches are old. My understanding is that the parts that are reasonable have been implemented and the parts that aren&#x27;t (plaintext modes) have been ignored.</p><p>If you need to transfer large files around a lot use what the supercomputing centers use: GridFTP.</p>",
              "points": null,
              "parent_id": 8829500,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8829781,
                  "created_at": "2015-01-03T03:12:45.000Z",
                  "created_at_i": 1420254765,
                  "type": "comment",
                  "author": "lotsofmangos",
                  "title": null,
                  "url": null,
                  "text": "<p>I&#x27;d never even heard of GridFtp before. It looks pretty good. I&#x27;ve taken to not installing ftp on servers at all and just using ssh, but this might make me think again.</p>",
                  "points": null,
                  "parent_id": 8829731,
                  "story_id": 8828543,
                  "children": [],
                  "options": []
                }
              ],
              "options": []
            }
          ],
          "options": []
        },
        {
          "id": 8830427,
          "created_at": "2015-01-03T10:24:33.000Z",
          "created_at_i": 1420280673,
          "type": "comment",
          "author": "jokoon",
          "title": null,
          "url": null,
          "text": "<p>A solution would be to have some kind of protocol revolving around ssh, that would open one ssh session, and use other socket connections to do other stuff, like ftp, scp, wget, etc...</p><p>Would be pretty hairy though. Maybe another solution would be to have some new socket standard that can achieve what both UDP and TCP do, at some kind of variable window size...</p>",
          "points": null,
          "parent_id": 8829208,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8828699,
      "created_at": "2015-01-02T21:38:35.000Z",
      "created_at_i": 1420234715,
      "type": "comment",
      "author": "forfengeligfaen",
      "title": null,
      "url": null,
      "text": "<p>If you believe Jacob Appelbaum, we probably should not be using SSH for anything <a href=\"http://media.ccc.de/browse/congress/2014/31c3_-_6258_-_en_-_saal_1_-_201412282030_-_reconstructing_narratives_-_jacob_-_laura_poitras.html#video\" rel=\"nofollow\">http:&#x2F;&#x2F;media.ccc.de&#x2F;browse&#x2F;congress&#x2F;2014&#x2F;31c3_-_6258_-_en_-_...</a></p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8828932,
          "created_at": "2015-01-02T22:29:32.000Z",
          "created_at_i": 1420237772,
          "type": "comment",
          "author": "iuguy",
          "title": null,
          "url": null,
          "text": "<p>If you believe that being able to sometimes compromise some implementations of SSH under some circumstances means that you shouldn&#x27;t use SSH for anything then maybe you shouldn&#x27;t use the Internet at all.</p>",
          "points": null,
          "parent_id": 8828699,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8828708,
          "created_at": "2015-01-02T21:40:09.000Z",
          "created_at_i": 1420234809,
          "type": "comment",
          "author": "shazow",
          "title": null,
          "url": null,
          "text": "<p>Yea I was fairly confused about that announcement as I couldn&#x27;t find anything damning in the released docs about SSH. Do you have a reference to a specific document&#x2F;slide?</p>",
          "points": null,
          "parent_id": 8828699,
          "story_id": 8828543,
          "children": [
            {
              "id": 8829204,
              "created_at": "2015-01-02T23:41:21.000Z",
              "created_at_i": 1420242081,
              "type": "comment",
              "author": "marcosdumay",
              "title": null,
              "url": null,
              "text": "<p>The slides say they can sometimes decrypt SSH sessions (with ssmintm?) and catpute usernames and passwords from SSH servers.</p><p>The only unexpected thing I could see there was the lack of anything I didn&#x27;t know how to do...</p>",
              "points": null,
              "parent_id": 8828708,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8829237,
                  "created_at": "2015-01-02T23:51:21.000Z",
                  "created_at_i": 1420242681,
                  "type": "comment",
                  "author": "shazow",
                  "title": null,
                  "url": null,
                  "text": "<p>Yea, sshmitm makes sense also, but I wouldn&#x27;t consider that a 0day exactly. :P</p>",
                  "points": null,
                  "parent_id": 8829204,
                  "story_id": 8828543,
                  "children": [],
                  "options": []
                }
              ],
              "options": []
            },
            {
              "id": 8828740,
              "created_at": "2015-01-02T21:47:10.000Z",
              "created_at_i": 1420235230,
              "type": "comment",
              "author": "doctorshady",
              "title": null,
              "url": null,
              "text": "<p>Spiegel briefly touched on it too;</p><p>The NSA also has a program with which it claims it can sometimes decrypt the Secure Shell protocol (SSH). This is typically used by systems administrators to log into employees&#x27; computers remotely, largely for use in the infrastructure of businesses, core Internet routers and other similarly important systems. The NSA combines the data collected in this manner with other information to leverage access to important systems of interest.</p><p>(source: <a href=\"http://www.spiegel.de/international/germany/inside-the-nsa-s-war-on-internet-security-a-1010361.html\" rel=\"nofollow\">http:&#x2F;&#x2F;www.spiegel.de&#x2F;international&#x2F;germany&#x2F;inside-the-nsa-s...</a> )</p><p>Still incredibly vague. If they&#x27;re archiving all traffic in \nhopes of decrypting it some day though, it&#x27;s safe to say we should treat anything on the internet as the shiny side of one way glass.</p>",
              "points": null,
              "parent_id": 8828708,
              "story_id": 8828543,
              "children": [],
              "options": []
            },
            {
              "id": 8829078,
              "created_at": "2015-01-02T23:04:41.000Z",
              "created_at_i": 1420239881,
              "type": "comment",
              "author": "boracay",
              "title": null,
              "url": null,
              "text": "<p>My guess is that, at least, some router ssh implementation is insecure, possibly not by accident.</p><p>From the slides:</p><p>Page 19: &quot;SSH [...] Potentially recover user names and passwords&quot;\nPage 36: &quot;SSH - often have router configurations and user credentials [...]&quot;</p><p><a href=\"http://www.spiegel.de/media/media-35515.pdf\" rel=\"nofollow\">http:&#x2F;&#x2F;www.spiegel.de&#x2F;media&#x2F;media-35515.pdf</a></p>",
              "points": null,
              "parent_id": 8828708,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8829159,
                  "created_at": "2015-01-02T23:28:50.000Z",
                  "created_at_i": 1420241330,
                  "type": "comment",
                  "author": "shazow",
                  "title": null,
                  "url": null,
                  "text": "<p>Right, I saw that also. Is it referring to routers that happen to run SSH with default root&#x2F;admin passwords or something? I couldn&#x27;t find anything more concrete.</p>",
                  "points": null,
                  "parent_id": 8829078,
                  "story_id": 8828543,
                  "children": [],
                  "options": []
                }
              ],
              "options": []
            }
          ],
          "options": []
        },
        {
          "id": 8828971,
          "created_at": "2015-01-02T22:39:06.000Z",
          "created_at_i": 1420238346,
          "type": "comment",
          "author": "tlrobinson",
          "title": null,
          "url": null,
          "text": "<p>Relevant bit starts around 00:25:30</p><p>Also: <a href=\"https://twitter.com/ioerror/status/549327936361611264\" rel=\"nofollow\">https:&#x2F;&#x2F;twitter.com&#x2F;ioerror&#x2F;status&#x2F;549327936361611264</a></p>",
          "points": null,
          "parent_id": 8828699,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8828744,
          "created_at": "2015-01-02T21:47:44.000Z",
          "created_at_i": 1420235264,
          "type": "comment",
          "author": "falcolas",
          "title": null,
          "url": null,
          "text": "<p>This link 404&#x27;s for me; can you give a summary?</p>",
          "points": null,
          "parent_id": 8828699,
          "story_id": 8828543,
          "children": [
            {
              "id": 8828827,
              "created_at": "2015-01-02T22:05:49.000Z",
              "created_at_i": 1420236349,
              "type": "comment",
              "author": "forfengeligfaen",
              "title": null,
              "url": null,
              "text": "<p>That&#x27;s weird all the CCC videos were returning 404. The link appears to be working again now. It&#x27;s a few days since I watched but I inferred that SSH may have some undisclosed vulnerabilities.</p>",
              "points": null,
              "parent_id": 8828744,
              "story_id": 8828543,
              "children": [],
              "options": []
            }
          ],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8829212,
      "created_at": "2015-01-02T23:44:16.000Z",
      "created_at_i": 1420242256,
      "type": "comment",
      "author": "malkia",
      "title": null,
      "url": null,
      "text": "<p>Because schools and filtering software</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [],
      "options": []
    },
    {
      "id": 8829981,
      "created_at": "2015-01-03T05:06:14.000Z",
      "created_at_i": 1420261574,
      "type": "comment",
      "author": "codingthebeach",
      "title": null,
      "url": null,
      "text": "<p>It was difficult for me to get past the self-referential (and reverential) tone of the writing, which felt more like an advertisement for the author&#x27;s cleverness than a real discussion of the pros and cons of SSH, but it could be I woke up on the wrong side of the bed this morning.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [],
      "options": []
    },
    {
      "id": 8829214,
      "created_at": "2015-01-02T23:44:30.000Z",
      "created_at_i": 1420242270,
      "type": "comment",
      "author": "malkia",
      "title": null,
      "url": null,
      "text": "<p>Because schools and filtering software</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [],
      "options": []
    },
    {
      "id": 8829727,
      "created_at": "2015-01-03T02:49:24.000Z",
      "created_at_i": 1420253364,
      "type": "comment",
      "author": "peterwwillis",
      "title": null,
      "url": null,
      "text": "<p>What the...? Does the author just not know anything about SSH, or web browsers? Why would we use SSH for everything?</p><p>On top of the fact that they&#x27;re entirely different protocols and tools designed for entirely different purposes, browsers already support virtually everything SSH does. File transfers, authentication, client certs, multiplexing, key pinning, etc. There is no need to use SSH, and if you did, it would be slower, less secure, and generally more annoying than using the existing tools built into browsers.</p><p>I find it aggravating when users read the manual to some software and think they have discovered fire.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829893,
          "created_at": "2015-01-03T04:08:17.000Z",
          "created_at_i": 1420258097,
          "type": "comment",
          "author": "nickpresta",
          "title": null,
          "url": null,
          "text": "<p>What does this even mean?</p><p>Your browser supports those things because it implements protocols, like HTTP, TLS, and others. Your browser, or other tools, <i></i>could<i></i> support SSH, which I think is the point of the article.</p>",
          "points": null,
          "parent_id": 8829727,
          "story_id": 8828543,
          "children": [
            {
              "id": 8830364,
              "created_at": "2015-01-03T09:46:11.000Z",
              "created_at_i": 1420278371,
              "type": "comment",
              "author": "peterwwillis",
              "title": null,
              "url": null,
              "text": "<p>This means the author doesn&#x27;t know what they&#x27;re talking about, and is asking a dumb question.</p><p>The article&#x27;s title is &quot;Why aren&#x27;t we using SSH for everything?&quot;, not &quot;Why don&#x27;t browsers support SSH?&quot;. Both make no sense. It&#x27;s like asking why FTP clients don&#x27;t support Voice-over-IP.</p>",
              "points": null,
              "parent_id": 8829893,
              "story_id": 8828543,
              "children": [],
              "options": []
            }
          ],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8828960,
      "created_at": "2015-01-02T22:36:30.000Z",
      "created_at_i": 1420238190,
      "type": "comment",
      "author": "josephg",
      "title": null,
      "url": null,
      "text": "<p>Because SSH requires several seconds to initiate a session, even on a local LAN.</p><p>Does anyone know why this is the case? Its always baffled me.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829751,
          "created_at": "2015-01-03T02:59:17.000Z",
          "created_at_i": 1420253957,
          "type": "comment",
          "author": "jloughry",
          "title": null,
          "url": null,
          "text": "<p>A delay of several seconds every time when connecting via ssh is usually due to the remote host trying to look up your IP address and timing out. Change the remote host&#x27;s sshd_config file to include the line &quot;UseDNS no&quot; and connecting will be much quicker from then on.</p>",
          "points": null,
          "parent_id": 8828960,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8828991,
          "created_at": "2015-01-02T22:44:45.000Z",
          "created_at_i": 1420238685,
          "type": "comment",
          "author": "shazow",
          "title": null,
          "url": null,
          "text": "<p>It shouldn&#x27;t take much longer than an HTTPS connection, could be something with your server configuration or latency?</p>",
          "points": null,
          "parent_id": 8828960,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8828993,
          "created_at": "2015-01-02T22:44:48.000Z",
          "created_at_i": 1420238688,
          "type": "comment",
          "author": "scott_karana",
          "title": null,
          "url": null,
          "text": "<p>TLS also takes some time to initiate a session, but that hasn&#x27;t turned anyone off SPDY, since it supports concurrency from then onwards. :-)</p>",
          "points": null,
          "parent_id": 8828960,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8829020,
          "created_at": "2015-01-02T22:52:22.000Z",
          "created_at_i": 1420239142,
          "type": "comment",
          "author": "agwa",
          "title": null,
          "url": null,
          "text": "<p>Initiating an SSH connection takes 300ms on my LAN.</p><p>If it&#x27;s taking you several seconds then you probably have a reverse DNS problem.  (By default, the SSH server looks up the reverse DNS of the client when it connects.)</p>",
          "points": null,
          "parent_id": 8828960,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8829794,
          "created_at": "2015-01-03T03:18:45.000Z",
          "created_at_i": 1420255125,
          "type": "comment",
          "author": "peterwwillis",
          "title": null,
          "url": null,
          "text": "<p>There are many possible causes, but the biggest one is that openssl is notoriously slow at the handshake process. Changing the cipher, playing with DNS, turning off unnecessary features and tuning the order of operations can speed up the process. Otherwise, try a different ssh daemon that doesn&#x27;t use openssl.</p>",
          "points": null,
          "parent_id": 8828960,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8829051,
          "created_at": "2015-01-02T22:58:32.000Z",
          "created_at_i": 1420239512,
          "type": "comment",
          "author": "encoded",
          "title": null,
          "url": null,
          "text": "<p>It could also be your shell on the remote host taking it&#x27;s time. Some distros have gotten a bit bloated.</p>",
          "points": null,
          "parent_id": 8828960,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8829059,
          "created_at": "2015-01-02T22:59:39.000Z",
          "created_at_i": 1420239579,
          "type": "comment",
          "author": "geoffc",
          "title": null,
          "url": null,
          "text": "<p>Slow SSH session is usually a DNS issue.</p>",
          "points": null,
          "parent_id": 8828960,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8829079,
          "created_at": "2015-01-02T23:05:00.000Z",
          "created_at_i": 1420239900,
          "type": "comment",
          "author": "monksy",
          "title": null,
          "url": null,
          "text": "<p>Thats probably because the server is trying to lookup the client on DNS: \n<a href=\"http://www.turnkeylinux.org/blog/slow-ssh\" rel=\"nofollow\">http:&#x2F;&#x2F;www.turnkeylinux.org&#x2F;blog&#x2F;slow-ssh</a></p>",
          "points": null,
          "parent_id": 8828960,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8829095,
          "created_at": "2015-01-02T23:08:22.000Z",
          "created_at_i": 1420240102,
          "type": "comment",
          "author": "_JamesA_",
          "title": null,
          "url": null,
          "text": "<p>You may be experiencing the GSSAPIAuthentication delay.</p><p>The -v flag will help identify where the connection is being delayed.</p><p>ssh -v [user@]hostname [command]</p><p>Try connecting with the option disabled.</p><p>ssh -o GSSAPIAuthentication=no [user@]hostname [command]</p><p>You can globally disable GSSAPIAuthentication in ssh config:</p><p>echo &#x27;GSSAPIAuthentication=no&#x27; &gt;&gt; ~&#x2F;.ssh&#x2F;config</p>",
          "points": null,
          "parent_id": 8828960,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8829496,
      "created_at": "2015-01-03T01:06:06.000Z",
      "created_at_i": 1420247166,
      "type": "comment",
      "author": "kgilpin",
      "title": null,
      "url": null,
      "text": "<p>Something else worth mentioning; SSH has built-in support for single-sign-on, via SSH agent forwarding. As long as my public key is available anywhere (and that&#x27;s exactly what it&#x27;s designed for), then I can be authenticated by any system, anywhere. Thus, a problem which is so vexing in so many other scenarios, is very cleanly addressed by SSH.</p><p>Fundamentally, a password is a shared secret. So you send your password to a server, you are trusting that server not to lose or misplace it. In contrast, an SSH public key doesn&#x27;t require nearly such careful management.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829796,
          "created_at": "2015-01-03T03:19:32.000Z",
          "created_at_i": 1420255172,
          "type": "comment",
          "author": "dsl",
          "title": null,
          "url": null,
          "text": "<p>SSH agent forwarding is extremely risky. Anyone with appropriate permissions (legitimate or illegitimately gained) on a machine you have connected to can use your credentials to open a new connection to any machine you have access to.</p>",
          "points": null,
          "parent_id": 8829496,
          "story_id": 8828543,
          "children": [
            {
              "id": 8830146,
              "created_at": "2015-01-03T07:13:22.000Z",
              "created_at_i": 1420269202,
              "type": "comment",
              "author": "Sevrene",
              "title": null,
              "url": null,
              "text": "<p>This is a very valid issue. I have been using ssh-ident* for this reason. You can add keys to ~&#x2F;.ssh&#x2F;identities and it&#x27;ll start up an agent for each identity when you use it. This means if the machine you connect to is indeed stealing your keys, at least they only have one key, not all the keys you&#x27;re currently using.</p><p>There are some caveats; such as having to specify what domains&#x2F;IPs should use which identity, and also having to alias ssh to get it to work nicely. So it&#x27;s not perfect, but it does solve the problem you talked about.</p><p>* <a href=\"https://github.com/ccontavalli/ssh-ident\" rel=\"nofollow\">https:&#x2F;&#x2F;github.com&#x2F;ccontavalli&#x2F;ssh-ident</a></p>",
              "points": null,
              "parent_id": 8829796,
              "story_id": 8828543,
              "children": [],
              "options": []
            },
            {
              "id": 8830164,
              "created_at": "2015-01-03T07:21:52.000Z",
              "created_at_i": 1420269712,
              "type": "comment",
              "author": "btown",
              "title": null,
              "url": null,
              "text": "<p>&gt; on a machine you have connected to</p><p>Don&#x27;t you have to be actively connected to the machine for this to work? i.e. the server I haven&#x27;t connected to for a few months has no way of opening connections on my behalf at this point, right?</p>",
              "points": null,
              "parent_id": 8829796,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8830236,
                  "created_at": "2015-01-03T08:20:32.000Z",
                  "created_at_i": 1420273232,
                  "type": "comment",
                  "author": "dsl",
                  "title": null,
                  "url": null,
                  "text": "<p>You are correct. That is what I meant, but I should have said &quot;a machine you are actively connected to.&quot;</p>",
                  "points": null,
                  "parent_id": 8830164,
                  "story_id": 8828543,
                  "children": [],
                  "options": []
                },
                {
                  "id": 8830175,
                  "created_at": "2015-01-03T07:28:23.000Z",
                  "created_at_i": 1420270103,
                  "type": "comment",
                  "author": "beagle3",
                  "title": null,
                  "url": null,
                  "text": "<p>You are right.</p><p>But once you connect, a year later - in those 3 minutes before you disconnect, the attacker might have authenticated as yourself to 100 other machines -- and appended their own key to .ssh&#x2F;authorized_keys on these machines, so that the compromise no longer needs you to be connected.</p>",
                  "points": null,
                  "parent_id": 8830164,
                  "story_id": 8828543,
                  "children": [],
                  "options": []
                }
              ],
              "options": []
            }
          ],
          "options": []
        },
        {
          "id": 8851421,
          "created_at": "2015-01-07T16:29:08.000Z",
          "created_at_i": 1420648148,
          "type": "comment",
          "author": "jamiesonbecker",
          "title": null,
          "url": null,
          "text": "<p>Agreed! On nomenclature - you&#x27;re saying ssh agent forwarding, but the functionality you describe is actually simply public keys on remote machines. The only time you need to forward is when you&#x27;re logging into another machine <i>through</i> a third machine -- and as other commenters have pointed out, you&#x27;re extending your trust model to that third machine.</p>",
          "points": null,
          "parent_id": 8829496,
          "story_id": 8828543,
          "children": [
            {
              "id": 8852532,
              "created_at": "2015-01-07T19:05:48.000Z",
              "created_at_i": 1420657548,
              "type": "comment",
              "author": "jamiesonbecker",
              "title": null,
              "url": null,
              "text": "<p>(blatant plug) Userify can distribute those public keys and reduce (but not eliminate) the cases where agent forwarding is required.</p>",
              "points": null,
              "parent_id": 8851421,
              "story_id": 8828543,
              "children": [],
              "options": []
            }
          ],
          "options": []
        },
        {
          "id": 8829674,
          "created_at": "2015-01-03T02:24:17.000Z",
          "created_at_i": 1420251857,
          "type": "comment",
          "author": "dllthomas",
          "title": null,
          "url": null,
          "text": "<p>Of course, you need to trust the remote systems...</p>",
          "points": null,
          "parent_id": 8829496,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8828743,
      "created_at": "2015-01-02T21:47:29.000Z",
      "created_at_i": 1420235249,
      "type": "comment",
      "author": "falcolas",
      "title": null,
      "url": null,
      "text": "<p>Neat trick, if you&#x27;re so inclined to use such tricks:</p><p><pre><code>    $ cat .ssh&#x2F;authorized_keys\n    command=&quot;tmux new-session -A -s base&quot; ssh-rsa [...]\n</code></pre>\nAutomatically creates or joins a tmux session named base, and disconnects the SSH session when you disconnect from the  tmux session.</p><p>So, yeah, why don&#x27;t we use SSH for more?</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8828872,
          "created_at": "2015-01-02T22:15:17.000Z",
          "created_at_i": 1420236917,
          "type": "comment",
          "author": "XorNot",
          "title": null,
          "url": null,
          "text": "<p>It&#x27;s multiplexing support uses static sized windows, which even on modern LANs means you usually only see 3-5 mb&#x2F;s transfer rates where you should be able to - even with encryption overhead - achieve almost gigabit NIC speed.</p><p>Fortunately the HPN-SSH patches exist to solve this problem - but I really want to know why their&#x27;s so much resistance to adding them upstream.</p>",
          "points": null,
          "parent_id": 8828743,
          "story_id": 8828543,
          "children": [
            {
              "id": 8829069,
              "created_at": "2015-01-02T23:02:47.000Z",
              "created_at_i": 1420239767,
              "type": "comment",
              "author": "falcolas",
              "title": null,
              "url": null,
              "text": "<p>Good answer.</p><p>If I had to guess, the reason would be threefold. One, it requires tuning on both ends of the connection to make the high speed transfers work well (and it looks like it will interfere with multiplexing as well, the reason the fixed window exists in the first place).</p><p>Two, the &quot;none&quot; cypher which it enables completely undermines the security of the secure shell. The cypher renegotiation after authentication would make it trivial to MITM the connection. You could probably disable the cypher after every transfer by restarting the sshd daemon, but that&#x27;s being a bit optimistic.</p><p>Finally, there is no visible license on the patchfile, or on the homepage. The project&#x27;s sourceforge page lists it as being BSD licensed, but there is no indication from the author that this is the case.</p><p><a href=\"http://www.psc.edu/index.php/hpn-ssh\" rel=\"nofollow\">http:&#x2F;&#x2F;www.psc.edu&#x2F;index.php&#x2F;hpn-ssh</a></p>",
              "points": null,
              "parent_id": 8828872,
              "story_id": 8828543,
              "children": [],
              "options": []
            },
            {
              "id": 8829339,
              "created_at": "2015-01-03T00:23:57.000Z",
              "created_at_i": 1420244637,
              "type": "comment",
              "author": "acdha",
              "title": null,
              "url": null,
              "text": "<p>&gt; which even on modern LANs means you usually only see 3-5 mb&#x2F;s transfer rates where you should be able to</p><p>On a WAN, sure, but even a mid-2000s gigabit LAN could routinely hit 800+Mb with scp as long as you had tuned the underlying TCP stack (Linux was poor and OS X&#x2F;*BSD worse in that era) and weren&#x27;t using something slow like 3-DES.</p><p>&gt; Fortunately the HPN-SSH patches exist to solve this problem - but I really want to know why their&#x27;s so much resistance to adding them upstream.</p><p>Looking at the patches, I&#x27;d be surprised if the problem wasn&#x27;t the fact that they change other things with security or reliability implications. Seeing something like “Dynamic Window and ability to use NONE encryption” suggests that it&#x27;d be better to break it up into some smaller generally-useful patches and a separate patch for people in controlled environments who need as much performance as possible.</p>",
              "points": null,
              "parent_id": 8828872,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8830444,
                  "created_at": "2015-01-03T10:35:19.000Z",
                  "created_at_i": 1420281319,
                  "type": "comment",
                  "author": "XorNot",
                  "title": null,
                  "url": null,
                  "text": "<p>It is broken up into separate patches as well (scroll down the homepage).</p><p>But the NONE encryption makes a lot of sense in a &quot;use SSH more&quot; perspective. When you need to move piles of totally non-private data, but want to use a secure authentication mechanism (and message authenticity system) for issuing shell commands...</p>",
                  "points": null,
                  "parent_id": 8829339,
                  "story_id": 8828543,
                  "children": [],
                  "options": []
                }
              ],
              "options": []
            },
            {
              "id": 8830188,
              "created_at": "2015-01-03T07:41:24.000Z",
              "created_at_i": 1420270884,
              "type": "comment",
              "author": "beagle3",
              "title": null,
              "url": null,
              "text": "<p>&gt; means you usually only see 3-5 mb&#x2F;s</p><p>I see over 50Mb&#x2F;s (=~400Mb&#x2F;s) on a gb network without any attempts at optimization, between a stupid Atom machine and a fast machine, both running Linux (which other responders claim is unoptimized). When did you last test?</p><p>Also, this might be the speed of the disks at one of the ends - I&#x27;ve never even bothered to check, because making it faster doesn&#x27;t make a difference for me.</p>",
              "points": null,
              "parent_id": 8828872,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8830452,
                  "created_at": "2015-01-03T10:42:42.000Z",
                  "created_at_i": 1420281762,
                  "type": "comment",
                  "author": "XorNot",
                  "title": null,
                  "url": null,
                  "text": "<p>My experience with a 20-disk RAIDZ3 machine and my desktops has always been that SSH tops out at 3-5mb&#x2F;s, which was why I could never understand the prevalence of advice to &quot;just do your disk backup over SSH&quot;. The same disks Samba between Windows and Linux achieve 70-90 mb&#x2F;s, and switching to HPN-SSH (in either Cygwin or Linux) gives me the same sort of speeds with the AES cipher or none cipher.</p><p>These are not slow machines - i7&#x27;s and i5&#x27;s.</p><p>I&#x27;ve never been able to find a compelling explanation as to how this could be when other network protocols can manage it just fine.</p>",
                  "points": null,
                  "parent_id": 8830188,
                  "story_id": 8828543,
                  "children": [
                    {
                      "id": 8830686,
                      "created_at": "2015-01-03T13:21:31.000Z",
                      "created_at_i": 1420291291,
                      "type": "comment",
                      "author": "beagle3",
                      "title": null,
                      "url": null,
                      "text": "<p>(I wanted to write 50MB =~ 400Mb earlier).</p><p>Something in your setup is broken. I get 50MB&#x2F;s speeds when Linux is on both ends, and even when Cygwin is on one end.</p>",
                      "points": null,
                      "parent_id": 8830452,
                      "story_id": 8828543,
                      "children": [],
                      "options": []
                    }
                  ],
                  "options": []
                }
              ],
              "options": []
            }
          ],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8828749,
      "created_at": "2015-01-02T21:49:09.000Z",
      "created_at_i": 1420235349,
      "type": "comment",
      "author": "higherpurpose",
      "title": null,
      "url": null,
      "text": "<p>I read somewhere that SSH can be MITM&#x27;ed by a global adversary on the first visit (before it establishes the secure connection). Is that true?</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8828984,
          "created_at": "2015-01-02T22:42:35.000Z",
          "created_at_i": 1420238555,
          "type": "comment",
          "author": "richardwhiuk",
          "title": null,
          "url": null,
          "text": "<p>You can fix this with SSHFP records and DNSSEC.</p>",
          "points": null,
          "parent_id": 8828749,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8828776,
          "created_at": "2015-01-02T21:53:40.000Z",
          "created_at_i": 1420235620,
          "type": "comment",
          "author": "forgottenpass",
          "title": null,
          "url": null,
          "text": "<p>It can be MITM-ed on any connection. Whether you approve the host fingerprint is, and has always been, entirely up to you.</p>",
          "points": null,
          "parent_id": 8828749,
          "story_id": 8828543,
          "children": [
            {
              "id": 8828813,
              "created_at": "2015-01-02T22:00:14.000Z",
              "created_at_i": 1420236014,
              "type": "comment",
              "author": "singlow",
              "title": null,
              "url": null,
              "text": "<p>Default configuration on all of the system&#x27;s I have used (mostly Ubuntu) is that you cannot just approve it. You have to actually remove it from the known hosts table, in effect starting a new &quot;first session&quot;. You can set it to allow override without deleting the entry though and some system may have this as the default.</p>",
              "points": null,
              "parent_id": 8828776,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8828863,
                  "created_at": "2015-01-02T22:13:38.000Z",
                  "created_at_i": 1420236818,
                  "type": "comment",
                  "author": "XorNot",
                  "title": null,
                  "url": null,
                  "text": "<p>Which is still irrelevant: in the default configuration the server I&#x27;m connecting to is <i>probably</i> not compromised, I&#x27;ve probably just done something to change keys or routes or names or whatever.</p><p>The problem is further back: the default configuration doesn&#x27;t make it easy to avoid these problems in the first place. If I&#x27;m logged in and do something which will change SSH host keys or the like, then the default needs to provide a way for me to make that information easily available globally to other clients which might be aware of this system.</p>",
                  "points": null,
                  "parent_id": 8828813,
                  "story_id": 8828543,
                  "children": [
                    {
                      "id": 8829404,
                      "created_at": "2015-01-03T00:40:53.000Z",
                      "created_at_i": 1420245653,
                      "type": "comment",
                      "author": "ytjohn",
                      "title": null,
                      "url": null,
                      "text": "<p>I don&#x27;t even know of a client that implements this, but SSFP DNS records are the way to do this.  <a href=\"http://www.openssh.com/txt/rfc4255.txt\" rel=\"nofollow\">http:&#x2F;&#x2F;www.openssh.com&#x2F;txt&#x2F;rfc4255.txt</a></p><p>A security conscious client can manually verify a host&#x27;s SSHFP with a dig record. You&#x27;ll also want DNSSEC in place to ensure the dns hasn&#x27;t been spoofed either.</p>",
                      "points": null,
                      "parent_id": 8828863,
                      "story_id": 8828543,
                      "children": [],
                      "options": []
                    }
                  ],
                  "options": []
                }
              ],
              "options": []
            }
          ],
          "options": []
        },
        {
          "id": 8828786,
          "created_at": "2015-01-02T21:55:48.000Z",
          "created_at_i": 1420235748,
          "type": "comment",
          "author": "josephlord",
          "title": null,
          "url": null,
          "text": "<p>You get shown a key fingerprint at first connection. If you verify that with a secure source you can confirm the host you are communicating with.</p><p>Actually bothering and having a secure way to verify the fingerprint are exercises left to the reader.</p>",
          "points": null,
          "parent_id": 8828749,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8828789,
          "created_at": "2015-01-02T21:56:11.000Z",
          "created_at_i": 1420235771,
          "type": "comment",
          "author": "rchowe",
          "title": null,
          "url": null,
          "text": "<p>If you don&#x27;t actually check the fingerprint of the key matches the fingerprint you were given (in some other secure way, such as walking over to the server), then yes, a first-visit MITM is possible.</p>",
          "points": null,
          "parent_id": 8828749,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8828801,
          "created_at": "2015-01-02T21:58:07.000Z",
          "created_at_i": 1420235887,
          "type": "comment",
          "author": "singlow",
          "title": null,
          "url": null,
          "text": "<p>Yes. It does not have any centralized certificate system like HTTPS so unless you can manually verify the host&#x27;s public key, you will not know whether your first visit is being proxied. Of course, if the first one is proxied, so may subsequent ones, and you would only get a warning if the proxy was removed or if it&#x27;s key changed.</p>",
          "points": null,
          "parent_id": 8828749,
          "story_id": 8828543,
          "children": [
            {
              "id": 8828814,
              "created_at": "2015-01-02T22:01:14.000Z",
              "created_at_i": 1420236074,
              "type": "comment",
              "author": "shazow",
              "title": null,
              "url": null,
              "text": "<p>SSH supports CA-style key signing, and it also supports server fingerprint validation over DNSSEC (search for SSHFP DNS).</p><p>Unfortunately neither of these things are commonly used yet. Cloudflare is adding DNSSEC support soon, so hopefully that will change.</p>",
              "points": null,
              "parent_id": 8828801,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8829046,
                  "created_at": "2015-01-02T22:58:09.000Z",
                  "created_at_i": 1420239489,
                  "type": "comment",
                  "author": "agwa",
                  "title": null,
                  "url": null,
                  "text": "<p>It&#x27;s going to take a lot more than Cloudflare adding DNSSEC support to make SSHFP records viable.  Every system running an SSH client will need to run its own validating resolver.  If you leave validation to an upstream server you lose a significant amount of security.</p>",
                  "points": null,
                  "parent_id": 8828814,
                  "story_id": 8828543,
                  "children": [
                    {
                      "id": 8831978,
                      "created_at": "2015-01-03T20:57:50.000Z",
                      "created_at_i": 1420318670,
                      "type": "comment",
                      "author": "tptacek",
                      "title": null,
                      "url": null,
                      "text": "<p>Not to mention that if your adversary is the Global Adversary, DNSSEC is mostly useless.</p>",
                      "points": null,
                      "parent_id": 8829046,
                      "story_id": 8828543,
                      "children": [],
                      "options": []
                    }
                  ],
                  "options": []
                }
              ],
              "options": []
            }
          ],
          "options": []
        },
        {
          "id": 8829830,
          "created_at": "2015-01-03T03:31:51.000Z",
          "created_at_i": 1420255911,
          "type": "comment",
          "author": "peterwwillis",
          "title": null,
          "url": null,
          "text": "<p>Yes. SSH can be silently mitm&#x27;d on the first connection due to not having the host key cached.</p><p>On the other hand, every web browser that visits a site for the first time (unless it was pinned in the browser you downloaded) attempts HTTP before HTTPS, and is thus vulnerable to all sorts of attacks. All non-pinned HTTPS connections at any time can be mitm&#x27;d by a global adversary that generates a cert using a CA your browser trusts.</p><p>So technically SSL is much easier to catch being mitm&#x27;d, since you only need to worry about the first visit.</p>",
          "points": null,
          "parent_id": 8828749,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8830204,
          "created_at": "2015-01-03T07:54:30.000Z",
          "created_at_i": 1420271670,
          "type": "comment",
          "author": "beagle3",
          "title": null,
          "url": null,
          "text": "<p>Yes, but only on the first visit. In comparison, TLS could be MITM&#x27;ed in each individual visit by a localized adversary as long as their location includes a CA you trust - which appears to be the case for every intelligence agency out there.</p><p>Which threat model you prefer to deal with is up to you.</p>",
          "points": null,
          "parent_id": 8828749,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8828754,
      "created_at": "2015-01-02T21:49:57.000Z",
      "created_at_i": 1420235397,
      "type": "comment",
      "author": "forgottenpass",
      "title": null,
      "url": null,
      "text": "<p><i>Why aren’t we using SSH for everything?</i></p><p>Because &quot;Use X for everything&quot; is a terrible design decision? SSH uses flexible transport with some desirable features and may be underutilized in practice.</p><p>This question is starting to feel like people who want to staple every pie in the sky idea to the bitcoin blockchain because it too has a set of desirable properties.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8828795,
          "created_at": "2015-01-02T21:57:03.000Z",
          "created_at_i": 1420235823,
          "type": "comment",
          "author": "shazow",
          "title": null,
          "url": null,
          "text": "<p>You&#x27;re right, we probably shouldn&#x27;t use SSH to microwave our food.</p><p>But a lot of things where we use HTTP today, we could be using SSH if we had better library support. Some more ideas towards the end of the post.</p>",
          "points": null,
          "parent_id": 8828754,
          "story_id": 8828543,
          "children": [
            {
              "id": 8829191,
              "created_at": "2015-01-02T23:37:50.000Z",
              "created_at_i": 1420241870,
              "type": "comment",
              "author": "marcosdumay",
              "title": null,
              "url": null,
              "text": "<p>&gt; we could be using SSH if we had better library support</p><p>It&#x27;s the firewall rules, not library availability.</p>",
              "points": null,
              "parent_id": 8828795,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8829340,
                  "created_at": "2015-01-03T00:24:07.000Z",
                  "created_at_i": 1420244647,
                  "type": "comment",
                  "author": "Shish2k",
                  "title": null,
                  "url": null,
                  "text": "<p>Tunnel SSH over HTTP, best (worst) of both worlds :)</p><p>(I have tried this; it is both useful and terrible)</p>",
                  "points": null,
                  "parent_id": 8829191,
                  "story_id": 8828543,
                  "children": [
                    {
                      "id": 8829378,
                      "created_at": "2015-01-03T00:35:10.000Z",
                      "created_at_i": 1420245310,
                      "type": "comment",
                      "author": "bashinator",
                      "title": null,
                      "url": null,
                      "text": "<p>Or you could just run 'sshd' on port 443.</p>",
                      "points": null,
                      "parent_id": 8829340,
                      "story_id": 8828543,
                      "children": [
                        {
                          "id": 8829431,
                          "created_at": "2015-01-03T00:50:12.000Z",
                          "created_at_i": 1420246212,
                          "type": "comment",
                          "author": "dec0dedab0de",
                          "title": null,
                          "url": null,
                          "text": "<p>But this would fail on many proxies, and any firewall that was aware of the protocol.</p>",
                          "points": null,
                          "parent_id": 8829378,
                          "story_id": 8828543,
                          "children": [],
                          "options": []
                        }
                      ],
                      "options": []
                    }
                  ],
                  "options": []
                }
              ],
              "options": []
            },
            {
              "id": 8828958,
              "created_at": "2015-01-02T22:36:04.000Z",
              "created_at_i": 1420238164,
              "type": "comment",
              "author": "blfr",
              "title": null,
              "url": null,
              "text": "<p>Someone has to be running NetBSD on their microwave owen. I remember a toaster story.</p><p><a href=\"http://www.embeddedarm.com/software/arm-netbsd-toaster.php\" rel=\"nofollow\">http:&#x2F;&#x2F;www.embeddedarm.com&#x2F;software&#x2F;arm-netbsd-toaster.php</a></p>",
              "points": null,
              "parent_id": 8828795,
              "story_id": 8828543,
              "children": [],
              "options": []
            },
            {
              "id": 8829269,
              "created_at": "2015-01-03T00:02:37.000Z",
              "created_at_i": 1420243357,
              "type": "comment",
              "author": "scrollaway",
              "title": null,
              "url": null,
              "text": "<p>Indeed. And more generally, using some of the SSH concepts on the web such as public key authentication.</p>",
              "points": null,
              "parent_id": 8828795,
              "story_id": 8828543,
              "children": [],
              "options": []
            }
          ],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8829012,
      "created_at": "2015-01-02T22:49:52.000Z",
      "created_at_i": 1420238992,
      "type": "comment",
      "author": "sharpneli",
      "title": null,
      "url": null,
      "text": "<p>SSH can really be used for almost everything. It&#x27;s a different thing if it actually _should_ be used for everything.</p><p>My favourite one: <a href=\"http://en.wikipedia.org/wiki/SSHFS\" rel=\"nofollow\">http:&#x2F;&#x2F;en.wikipedia.org&#x2F;wiki&#x2F;SSHFS</a></p><p>Whenever I&#x27;m developing for my mobile phone I actually have the contents mounted on my desktop via sshfs as an actual filesystem.</p><p>What that means? Don&#x27;t bother with FTP servers. ssh access to your server is all you need.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [],
      "options": []
    },
    {
      "id": 8829292,
      "created_at": "2015-01-03T00:10:34.000Z",
      "created_at_i": 1420243834,
      "type": "comment",
      "author": "ashish01",
      "title": null,
      "url": null,
      "text": "<p>Would be great to do this for mosh - mosh-chat It improves over SSH by handling intermittent network connections.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829333,
          "created_at": "2015-01-03T00:22:43.000Z",
          "created_at_i": 1420244563,
          "type": "comment",
          "author": "Shish2k",
          "title": null,
          "url": null,
          "text": "<p>mosh isn&#x27;t a data stream tool like ssh though, it&#x27;s actually more like VNC -- the reason it performs better is that it sends snapshots of the terminal over UDP, allowing random packets to be dropped</p>",
          "points": null,
          "parent_id": 8829292,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8829294,
      "created_at": "2015-01-03T00:11:15.000Z",
      "created_at_i": 1420243875,
      "type": "comment",
      "author": "throwaway593",
      "title": null,
      "url": null,
      "text": "<p>&gt; Why aren’t we using SSH for everything?</p><p>Because it doesn&#x27;t support virtual hosts. And I can&#x27;t afford 30 IPs for my server.</p><p>Otherwise, it&#x27;s a great protocol.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829330,
          "created_at": "2015-01-03T00:21:54.000Z",
          "created_at_i": 1420244514,
          "type": "comment",
          "author": "TylerE",
          "title": null,
          "url": null,
          "text": "<p>Couldn&#x27;t you use the same IP and different ports? There&#x27;s nothing magic about port 22 - and in fact on production servers you should almost certainly change 22 to something else.</p>",
          "points": null,
          "parent_id": 8829294,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8829053,
      "created_at": "2015-01-02T22:58:52.000Z",
      "created_at_i": 1420239532,
      "type": "comment",
      "author": "ivan_ah",
      "title": null,
      "url": null,
      "text": "<p>Isn&#x27;t there a problem when you tunnel TCP over TCP with increasing window sizes (auto throttling mechanism meant to prevent packet fragmentation)?</p><p>Every time I&#x27;ve tried to keep a long-running ssh tunnel for  printing &#x2F; http, the connection degrades after a while. I&#x27;m sure there are some flags that can be set, but I thought this was the major show stopper for the &quot;everything over shh&quot; (since ssh uses TCP protocol)</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829098,
          "created_at": "2015-01-02T23:08:57.000Z",
          "created_at_i": 1420240137,
          "type": "comment",
          "author": "SilasX",
          "title": null,
          "url": null,
          "text": "<p>Have you tried the autossh utility? That&#x27;s what I use to keep up long-term tunnels.</p>",
          "points": null,
          "parent_id": 8829053,
          "story_id": 8828543,
          "children": [
            {
              "id": 8829197,
              "created_at": "2015-01-02T23:39:18.000Z",
              "created_at_i": 1420241958,
              "type": "comment",
              "author": "otterley",
              "title": null,
              "url": null,
              "text": "<p>The window size issue has to do with bandwidth, not session duration. See, e.g., <a href=\"http://www.psc.edu/index.php/hpn-ssh\" rel=\"nofollow\">http:&#x2F;&#x2F;www.psc.edu&#x2F;index.php&#x2F;hpn-ssh</a></p>",
              "points": null,
              "parent_id": 8829098,
              "story_id": 8828543,
              "children": [],
              "options": []
            }
          ],
          "options": []
        },
        {
          "id": 8830184,
          "created_at": "2015-01-03T07:35:17.000Z",
          "created_at_i": 1420270517,
          "type": "comment",
          "author": "beagle3",
          "title": null,
          "url": null,
          "text": "<p>Simple port tunnelling (-L and -R command line arguments) do not suffer from TCP over TCP. And if you want VPN-style usage, look into sshuttle. It is one way only, and that&#x27;s a good thing! (Most of the time you only need connections going one way and none going the other way). If you do want two way, either use two sshuttle connections (one each way) or OpenVPN.</p><p>The only thing about sshuttle that I&#x27;ve encountered that exposes it&#x27;s non-VPNness, is that the connection truly originate from the remote system - e.g. if you connect through sshuttle to your peer, the connection go from the peer to the peer on 127.0.0.1. That may or may not be a problem (e.g., logging is much less useful this way).</p><p>Try sshuttle. I&#x27;ve stopped using VPNs since I started.</p>",
          "points": null,
          "parent_id": 8829053,
          "story_id": 8828543,
          "children": [
            {
              "id": 8852540,
              "created_at": "2015-01-07T19:07:27.000Z",
              "created_at_i": 1420657647,
              "type": "comment",
              "author": "jamiesonbecker",
              "title": null,
              "url": null,
              "text": "<p>I have to agree that VPN&#x27;s are usually a bit too much; generally speaking, VPN&#x27;s are a chainsaw when a scalpel is more appropriate!</p>",
              "points": null,
              "parent_id": 8830184,
              "story_id": 8828543,
              "children": [],
              "options": []
            }
          ],
          "options": []
        },
        {
          "id": 8829419,
          "created_at": "2015-01-03T00:45:53.000Z",
          "created_at_i": 1420245953,
          "type": "comment",
          "author": "alexgartrell",
          "title": null,
          "url": null,
          "text": "<p>It depends how you set up the forwarding.</p><p>Casually looking at strings &#x2F;usr&#x2F;sbin&#x2F;sshd on my devserver reveals that there is a mode to create a tun device, which sets up a more traditional VPN-ish thing (basically, your computer sees a new interface and routes certain sections of IP space to that prefix.  sshd sucks in these packets and forwards them as is).  This results in two layers of congestion control, which (IIRC -- I&#x27;m not a TCP guy) Congestion Control Algorithms aren&#x27;t set up to handle very well.</p><p>The more traditional mechanism for tunneling though (the one most often referenced in the posts and things I&#x27;ve seen) is a layer 7 forwarder.  Basically it creates sockets to listen for connections on either ends on a specific port, recv&#x27;s data as a traditional network user, and then turns around and sends the data over the multiplexed transport to the end server.  It essentially strips away the congestion control on the first hop (because you&#x27;re just going to localhost anyway) and there&#x27;s no TCP&#x2F;IP packets in TCP&#x2F;IP packets thing going on anymore, just your application stream in their header transport thing.  It&#x27;s less flexible, but is good enough if you want to set up a simple proxy or something and doesn&#x27;t have the same pitfalls.</p><p>If you need the flexibility, something like OpenVPN might be good (though I haven&#x27;t used it myself).</p><p>And you might also want to look at stunnel as an ssl tunnel alternative that lets you have more control over what certs go where without archaic and dangerous &#x2F;etc&#x2F;ssh&#x2F;*_config meddling</p>",
          "points": null,
          "parent_id": 8829053,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8829063,
      "created_at": "2015-01-02T23:01:47.000Z",
      "created_at_i": 1420239707,
      "type": "comment",
      "author": "spydum",
      "title": null,
      "url": null,
      "text": "<p>Interesting idea, but isn&#x27;t ssh very sluggish when it comes to throughput and latency? Doesn&#x27;t it cost quite a bit more in CPU? I&#x27;d hate to make ssh the protocol replacement for http for a busy site.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829074,
          "created_at": "2015-01-02T23:03:50.000Z",
          "created_at_i": 1420239830,
          "type": "comment",
          "author": "monksy",
          "title": null,
          "url": null,
          "text": "<p>It depends on the clipher you&#x27;re using and the implementation&#x2F;hardware.</p>",
          "points": null,
          "parent_id": 8829063,
          "story_id": 8828543,
          "children": [
            {
              "id": 8829200,
              "created_at": "2015-01-02T23:39:52.000Z",
              "created_at_i": 1420241992,
              "type": "comment",
              "author": "otterley",
              "title": null,
              "url": null,
              "text": "<p>It depends more on the TCP window size; OpenSSH uses a small one by default and it cannot be changed.</p><p>See also <a href=\"http://www.psc.edu/index.php/hpn-ssh\" rel=\"nofollow\">http:&#x2F;&#x2F;www.psc.edu&#x2F;index.php&#x2F;hpn-ssh</a></p>",
              "points": null,
              "parent_id": 8829074,
              "story_id": 8828543,
              "children": [],
              "options": []
            }
          ],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8829114,
      "created_at": "2015-01-02T23:13:21.000Z",
      "created_at_i": 1420240401,
      "type": "comment",
      "author": "ezequiel-garzon",
      "title": null,
      "url": null,
      "text": "<p>What surprises me is the lack of total CLI control, potentially through SSH and ideally using keys, of your hosting provider&#x27;s control panel. Whether it&#x27;s CLI or some form of TUI, it&#x27;s bound to be faster and more convenient for many developers.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829167,
          "created_at": "2015-01-02T23:31:53.000Z",
          "created_at_i": 1420241513,
          "type": "comment",
          "author": "blfr",
          "title": null,
          "url": null,
          "text": "<p>Cloud providers nowadays usually have some sort of CLI clients. OpenStack comes with a full suite[1]. Google rolled out their own for Google Cloud[2]. And if there&#x27;s nothing official, there are often tools built by users available for interacting with the API.</p><p>Authentication is based on some sort of shared secret rather than keys though, yes.</p><p>[1] <a href=\"http://docs.openstack.org/user-guide/content/ch_cli.html\" rel=\"nofollow\">http:&#x2F;&#x2F;docs.openstack.org&#x2F;user-guide&#x2F;content&#x2F;ch_cli.html</a></p><p>[2] <a href=\"https://cloud.google.com/sdk/gcloud/\" rel=\"nofollow\">https:&#x2F;&#x2F;cloud.google.com&#x2F;sdk&#x2F;gcloud&#x2F;</a></p>",
          "points": null,
          "parent_id": 8829114,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8829647,
      "created_at": "2015-01-03T02:09:50.000Z",
      "created_at_i": 1420250990,
      "type": "comment",
      "author": "bondolo",
      "title": null,
      "url": null,
      "text": "<p>Key management is pretty primitive. It would be nice if SSH integrated better with PGP&#x2F;GPG. I recently spent too much time messing with monkeysphere, keychain, gpg-agent, gpgsm etc. trying to use GPG derived keys for SSH. While I could cobble something together I didn&#x27;t feel that it was a solution I could recommend to others as a general &quot;best practice&quot; because it involved too much installing and configuring and was still somewhat brittle (eg. providing key names to gpg-agent in my .bashrc file).</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8830094,
          "created_at": "2015-01-03T06:33:10.000Z",
          "created_at_i": 1420266790,
          "type": "comment",
          "author": "dredmorbius",
          "title": null,
          "url": null,
          "text": "<p>But SSH keys really don&#x27;t matter.</p><p>In PGP&#x2F;GPG, key persistence matters <i>because you&#x27;re using them to decrypt messages.  Long after content was created you may need the key to decrypt it.</p><p>For SSH, the key is only strictly necessary during the session.  Key </i>distribution* (of your public key to systems you need access to) is a bit of a pain, but between having your private key(s) where you need them, and authorized keys on various servers, there&#x27;s not all that much to worry about.  Host keys, perhaps, if you want to be rigorous about security.</p>",
          "points": null,
          "parent_id": 8829647,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8829670,
          "created_at": "2015-01-03T02:22:20.000Z",
          "created_at_i": 1420251740,
          "type": "comment",
          "author": "dllthomas",
          "title": null,
          "url": null,
          "text": "<p>I&#x27;d love to see a project to make this suck less.</p>",
          "points": null,
          "parent_id": 8829647,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8828626,
      "created_at": "2015-01-02T21:22:02.000Z",
      "created_at_i": 1420233722,
      "type": "comment",
      "author": "feld",
      "title": null,
      "url": null,
      "text": "<p>I wonder if his chat server has been hit by a botnet yet trying to ssh in and then sending tons of shell commands</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8828642,
          "created_at": "2015-01-02T21:24:56.000Z",
          "created_at_i": 1420233896,
          "type": "comment",
          "author": "4mnt",
          "title": null,
          "url": null,
          "text": "<p>That would be very interesting to watch (albeit a bit annoying for the people trying to chat)</p>",
          "points": null,
          "parent_id": 8828626,
          "story_id": 8828543,
          "children": [],
          "options": []
        },
        {
          "id": 8828655,
          "created_at": "2015-01-02T21:28:08.000Z",
          "created_at_i": 1420234088,
          "type": "comment",
          "author": "shazow",
          "title": null,
          "url": null,
          "text": "<p>Lots of people spamming&#x2F;flooding&#x2F;DoS&#x27;ing &quot;for fun&quot;, but no clueless bots stumbling in accidentally yet.</p><p>Btw previous HN thread about ssh-chat here: <a href=\"https://news.ycombinator.com/item?id=8743374\" rel=\"nofollow\">https:&#x2F;&#x2F;news.ycombinator.com&#x2F;item?id=8743374</a></p>",
          "points": null,
          "parent_id": 8828626,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8828633,
      "created_at": "2015-01-02T21:23:01.000Z",
      "created_at_i": 1420233781,
      "type": "comment",
      "author": "shazow",
      "title": null,
      "url": null,
      "text": "<p>Sorry the title has been revised to &quot;Why aren’t we using SSH for everything?&quot; in case mods see this.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [],
      "options": []
    },
    {
      "id": 8832746,
      "created_at": "2015-01-04T01:07:54.000Z",
      "created_at_i": 1420333674,
      "type": "comment",
      "author": "ramigb",
      "title": null,
      "url": null,
      "text": "<p>There is something about Finnish products that make them very strong, Nokia, SSH, Rovio.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [],
      "options": []
    },
    {
      "id": 8829935,
      "created_at": "2015-01-03T04:39:20.000Z",
      "created_at_i": 1420259960,
      "type": "comment",
      "author": "101914",
      "title": null,
      "url": null,
      "text": "<p>MOre interesting question: Why aren&#x27;t we using SSH-style authentication for everything?  And although they can be used with SSH, I do not mean certificates.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8830028,
          "created_at": "2015-01-03T05:42:42.000Z",
          "created_at_i": 1420263762,
          "type": "comment",
          "author": "iancarroll",
          "title": null,
          "url": null,
          "text": "<p>I don&#x27;t see why client side certificate authentication is excluded? You can easily self sign and create your own CA. There aren&#x27;t really any downsides. Putty with cryptoapi support exists.</p>",
          "points": null,
          "parent_id": 8829935,
          "story_id": 8828543,
          "children": [],
          "options": []
        }
      ],
      "options": []
    },
    {
      "id": 8829176,
      "created_at": "2015-01-02T23:34:27.000Z",
      "created_at_i": 1420241667,
      "type": "comment",
      "author": "sysk",
      "title": null,
      "url": null,
      "text": "<p>It&#x27;s not clear to me how SSH differs from SSL&#x2F;TLS conceptually. It seems to me both achieve similar goals (encrypted tunnel, client&#x2F;server authentication). Perhaps we should take the best bits of both protocols and create a new one? But then, I am reminded of <a href=\"http://xkcd.com/927/\" rel=\"nofollow\">http:&#x2F;&#x2F;xkcd.com&#x2F;927&#x2F;</a>.</p>",
      "points": null,
      "parent_id": 8828543,
      "story_id": 8828543,
      "children": [
        {
          "id": 8829855,
          "created_at": "2015-01-03T03:43:23.000Z",
          "created_at_i": 1420256603,
          "type": "comment",
          "author": "peterwwillis",
          "title": null,
          "url": null,
          "text": "<p>SSH is designed to serve a single service on a single host. It distributes its host key on the first connection and caches it indefinitely, assuming it will never change. SSH is designed with a limited set of protocol features, and everything else is kind of hacked on top of proprietary client&#x2F;server pairs. SSH is designed as a loose encrypted session (kind of like a pipe) for an application on a host.</p><p>TLS is designed to serve multiple services on multiple hosts. It depends on your browser trusting an intermediary host which validates the host key, so (in theory) the initial connection can&#x27;t be MITM&#x27;d, and so the key can change at any time or there can be multiple keys (which is needed for hosting multiple services on multiple hosts). TLS is designed to integrate tightly into an application.</p><p>When you compare the two protocols, TLS is clearly superior to SSH. But in terms of the features they support (tunneling, authentication, etc), it&#x27;s up to the server to add missing features outside of the protocol to provide for what the client wants to do.</p><p>For example, the SSH protocol basically provides an encrypted connection through which you can do whatever you want, similar to TLS. To do IP tunneling with SSH the application server activates extra functionality to connect the encrypted session to a driver which opens an IP tunnel. Or to authenticate your ssh session against a kerberos server, the ssh server does the actual kerberos authentication; the protocol just informs the client of what &#x27;basic&#x27; methods they can use, and the client tries to use one that works with the server&#x27;s methods.</p><p>Incidentally, TLS the protocol supports client certificate authentication, which provides similar functionality to SSH&#x27;s public keys. The HTTP protocol also does certificate pinning.</p>",
          "points": null,
          "parent_id": 8829176,
          "story_id": 8828543,
          "children": [
            {
              "id": 8831826,
              "created_at": "2015-01-03T20:11:37.000Z",
              "created_at_i": 1420315897,
              "type": "comment",
              "author": "onnoonno",
              "title": null,
              "url": null,
              "text": "<p>Maybe it would be nice then if TLS would somehow cache host keys, too? Like my browser caching the relationship:</p><p>www.bank.com is 1.2.3.4 with pub key XYZ</p><p>Or is this already implemented and I am too stupid to find it?</p>",
              "points": null,
              "parent_id": 8829855,
              "story_id": 8828543,
              "children": [
                {
                  "id": 8834472,
                  "created_at": "2015-01-04T14:25:07.000Z",
                  "created_at_i": 1420381507,
                  "type": "comment",
                  "author": "peterwwillis",
                  "title": null,
                  "url": null,
                  "text": "<p>Since this kind of goes against the point of using an intermediary to verify the host key, and multiple services and hosts make this much more difficult to support, it&#x27;s not built into TLS. But the application can add support for it. There are experimental web standards and methods for various OSes&#x2F;applications here <a href=\"https://www.owasp.org/index.php/Certificate_and_Public_Key_Pinning#Examples_of_Pinning\" rel=\"nofollow\">https:&#x2F;&#x2F;www.owasp.org&#x2F;index.php&#x2F;Certificate_and_Public_Key_P...</a></p>",
                  "points": null,
                  "parent_id": 8831826,
                  "story_id": 8828543,
                  "children": [],
                  "options": []
                }
              ],
              "options": []
            }
          ],
          "options": []
        }
      ],
      "options": []
    }
  ],
  "options": []
}
`
