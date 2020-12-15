# go-sro-fileutils

Provides utilities to read and parse files of the game Silkroad Online, written in Golang. This repository is mainly
used by [go-sro-agent-server](https://github.com/ferdoran/go-sro-agent-server)

## Features

- PK2 Reader
- Navmesh Parser

## Acknowledgement

As the development was not a single person's effort, I want to
thank [DaxterSoul](https://www.elitepvpers.com/forum/members/1084164-daxtersoul.html)
for sharing his wide knowledge on the game and its peculiarities.

Without his packet and file structure documentation this would not have been possible.

## Additional Projects

As this is just a framework, there are also projects taking this framework into use:

- [go-sro-framework](https://github.com/ferdoran/go-sro-framework)
- [go-sro-gateway-server](https://github.com/ferdoran/go-sro-gateway-server)
- [go-sro-agent-server](https://github.com/ferdoran/go-sro-agent-server)
## Contribution

If you want to engage in the development, you are free to so. Simply fork this project and submit your changes via Pull
Requests.

Providing a usable documentation is still a big item on my TODO list for this project. Any help on this is highly
appreciated.

If you have any more questions, feel free to contact me.

## Licensing

go-sro-fileutils is licensed under the DON'T BE A DICK PUBLIC LICENSE. See [LICENSE](LICENSE) for the full license text

---

## Backup / Old Stuff

### Item Stats Abbreviations

| Abbreviation | Description |
|--------------|-------------|
| ER | Evasion Rate / Parry Rate |
| PAR | Physical Absorbtion Rate |
| MAR | Magical Absorbtion Rate |
| PD | Physical Defense | 
| MD | Magical Defense |
| Dur | Durability |

### Degree and Tier
```csharp
public int DegreeOffset => ItemClass - (3 * ((ItemClass - 1) / 3)) - 1; //sro_client.sub_8BA6E0
public int Degree => ((ItemClass - 1) / 3) + 1;
```

### Navigation mesh

The navigation mesh (navmesh) contains information about which regions are walkable and which are not. First of all
let's look at which files exists and how they are related to each other.

The navmesh data is contained inside the `Data.pk2/navmesh` folder. It includes different file types:

* `mapinfo.mfo` which includes map information. Especially which regions exist / are used
* `nv_XXYY.nvm` which are the navmesh files. One nav mesh file describes one region
* `.dat` which are nav data for the AI
* `.ifo` which include some more information regarding objects on the map

A region is a field with a size of 1920 x 1920 units. Usually a region has a regular grid, that consists of squares with
a minimum size of 20 x 20 units per square. So in total there can be 96 * 96 squares inside a region.

There can also be just 16:
```ascii
╔════════════╦════════════╦════════════╦════════════╗
║            ║            ║            ║            ║
║    0,0     ║    1,0     ║    2,0     ║    3,0     ║
║            ║            ║            ║            ║
║            ║            ║            ║            ║
╠════════════╬════════════╬════════════╬════════════╣
║            ║            ║            ║            ║
║    0,1     ║    1,1     ║    2,1     ║    3,1     ║
║            ║            ║            ║            ║
║            ║            ║            ║            ║
╠════════════╬════════════╬════════════╬════════════╣
║            ║            ║            ║            ║
║    0,2     ║    1,2     ║    2,2     ║    3,2     ║
║            ║            ║            ║            ║
║            ║            ║            ║            ║
╠════════════╬════════════╬════════════╬════════════╣
║            ║            ║            ║            ║
║    0,3     ║    1,3     ║    2,3     ║    3,3     ║
║            ║            ║            ║            ║
║            ║            ║            ║            ║
╚════════════╩════════════╩════════════╩════════════╝
```

However, the fields inside of a region do not have to be squares. Usually they are just rectangles.

Each field inside of a region is called `cell`. Each cell is connected to other `cell`s inside the region. This
connection is called `edge`. This applies to regions as well. So the `edge`s inside a region are called `internal edge`
s.
`Edge`s that connect regions are called `global edge`s.

Each region is identified by its coordinate on the world map. As you probably have seen in the `nv_XXYY.nvm` pattern,
the size of the world map is limited to 2 bytes. 1 byte for the x coordinate and 1 byte for the y or z coordinate. So
its theoretical max size is 256 * 256. However, this is not 100% true, since 1 bit is used for the dungeon flag. So the
real size is 2<sup>8</sup> x 2<sup>7</sup>
or 256 * 128. So theoretically there can be 32,768 regions on the world map, but
not all are used. The vanilla SRO files have ~3300 regions.
